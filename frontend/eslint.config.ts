import eslint from '@eslint/js'
import tseslint, { type ConfigArray } from 'typescript-eslint'
import pluginVue from 'eslint-plugin-vue'
import oxlint from 'eslint-plugin-oxlint'
import configPrettier from 'eslint-config-prettier'
import { noLegacyVBtnPropsRule } from './eslint/rules/noLegacyVBtnPropsRule'

const temporalNoDateRestrictions = [
  {
    selector: "NewExpression[callee.name='Date']",
    message:
      'Use Temporal instead of constructing Date directly outside explicit native-Date boundaries.',
  },
  {
    selector: "CallExpression[callee.name='Date']",
    message:
      'Use Temporal instead of calling Date directly outside explicit native-Date boundaries.',
  },
  {
    selector: "CallExpression[callee.object.name='Date'][callee.property.name='now']",
    message: 'Use Temporal.Now instead of Date.now().',
  },
  {
    selector: "CallExpression[callee.object.name='Date'][callee.property.name='parse']",
    message: 'Use Temporal parsing instead of Date.parse().',
  },
  {
    selector: "CallExpression[callee.object.name='Date'][callee.property.name='UTC']",
    message: 'Use Temporal instead of Date.UTC().',
  },
  {
    selector: "TSTypeReference > Identifier[name='Date']",
    message:
      'Use Temporal types instead of the native Date type outside explicit native-Date boundaries.',
  },
]

const e2eLocatorHygieneRestrictions = [
  {
    selector: "CallExpression[callee.property.name='waitForSelector']",
    message: 'Use page.locator(...).waitFor() instead of page.waitForSelector().',
  },
  {
    selector: "CallExpression[callee.property.name='waitForTimeout']",
    message:
      'Prefer web-first expect() or locator.waitFor(); for an unavoidable fixed settle delay use settlePage from e2e/helpers/settle.',
  },
  {
    selector: "CallExpression[callee.property.name='$']",
    message:
      'Use page.locator() instead of page.$; locators auto-wait and fail loudly on ambiguous matches.',
  },
  {
    selector: "CallExpression[callee.property.name='$$']",
    message:
      'Use page.locator() with toHaveCount() instead of page.$$; locators auto-wait and fail loudly on ambiguous matches.',
  },
  {
    selector: "CallExpression[callee.property.name='pause']",
    message:
      'Use playwright --debug or --ui for interactive debugging; committed tests must not pause.',
  },
]

const config: ConfigArray = [
  {
    ignores: ['dist/**', 'node_modules/**', 'public/**', '*.config.cjs'],
  },

  // Core JS rules
  eslint.configs.recommended,
  {
    linterOptions: {
      reportUnusedDisableDirectives: 'off',
    },
    plugins: {
      '@typescript-eslint': tseslint.plugin,
    },
    rules: {
      // TypeScript and vue-tsc know browser, Node, and DOM globals; ESLint does not.
      'no-undef': 'off',
    },
  },

  // Vue 3 recommended rules
  ...pluginVue.configs['flat/recommended'],

  // Parse TypeScript without constructing a project-wide TypeScript program.
  {
    files: ['**/*.{ts,tsx}'],
    languageOptions: {
      parser: tseslint.parser,
    },
  },

  // vue-eslint-parser handles .vue files; tseslint.parser handles their <script> blocks
  {
    files: ['**/*.vue'],
    plugins: {
      local: {
        rules: {
          'no-legacy-v-btn-props': noLegacyVBtnPropsRule,
        },
      },
      '@typescript-eslint': tseslint.plugin,
    },
    languageOptions: {
      parserOptions: {
        parser: tseslint.parser,
        extraFileExtensions: ['.vue'],
      },
    },
    rules: {
      'local/no-legacy-v-btn-props': 'error',
      // Oxlint skips Vue SFCs for this rule because template usage is ambiguous.
      '@typescript-eslint/consistent-type-imports': [
        'error',
        { prefer: 'type-imports', fixStyle: 'inline-type-imports' },
      ],
    },
  },

  // Disable rules delegated to Oxlint's native implementation.
  ...oxlint.configs['flat/recommended'],
  {
    rules: {
      'no-unused-vars': 'off',
      // Keep frontend runtime and tests on the Temporal model; native Date is allowed
      // only in explicit adapter files that integrate with APIs requiring Date objects.
      'no-restricted-syntax': ['error', ...temporalNoDateRestrictions],
    },
  },

  // E2E locator hygiene: auto-waiting locators over raw page queries, and no
  // fixed sleeps outside the single documented settle helper.
  {
    files: ['e2e/**/*.ts'],
    rules: {
      'no-restricted-syntax': [
        'error',
        ...temporalNoDateRestrictions,
        ...e2eLocatorHygieneRestrictions,
      ],
    },
  },
  {
    files: ['e2e/helpers/settle.ts'],
    rules: {
      'no-restricted-syntax': ['error', ...temporalNoDateRestrictions],
    },
  },
  // Diagnostic repro entrypoints run outside test:e2e and intentionally use raw page APIs.
  {
    files: ['e2e/repro/**/*.ts'],
    rules: {
      'no-restricted-syntax': ['error', ...temporalNoDateRestrictions],
    },
  },

  {
    files: [
      'src/components/DatePicker.vue',
      'src/components/DatePicker.test.ts',
      'src/components/DatePicker.spec.ts',
    ],
    rules: {
      'no-restricted-syntax': 'off',
    },
  },

  // OpenAPI Typescript emits index signatures for map schemas.
  {
    files: ['src/types/api.ts'],
    rules: {
      '@typescript-eslint/consistent-indexed-object-style': 'off',
    },
  },

  // Must be last — disables formatting rules that conflict with Prettier
  configPrettier,
]

export default config
