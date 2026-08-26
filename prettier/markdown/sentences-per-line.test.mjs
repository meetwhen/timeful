import { describe, expect, it } from 'vitest'
import * as prettier from 'prettier'
import * as sentencesPerLine from './sentences-per-line.js'

function format(markdown, options = {}) {
  return prettier.format(markdown, {
    parser: 'markdown',
    plugins: [sentencesPerLine],
    ...options,
  })
}

describe('table-safe sentences-per-line formatter', () => {
  it('splits prose sentences onto separate physical lines', async () => {
    await expect(format('The first sentence ends here. The second sentence follows.\n')).resolves.toBe(
      'The first sentence ends here.\nThe second sentence follows.\n',
    )
  })

  it('keeps a multi-sentence table cell on its row', async () => {
    const formatted = await format(
      '| Scenario | Requirement |\n| --- | --- |\n| Response | The system shall issue access only after approval. It shall reject every invalid state. |\n',
    )

    expect(formatted).toContain(
      '| Response | The system shall issue access only after approval. It shall reject every invalid state. |',
    )
    expect(formatted.split('\n').filter((line) => line.startsWith('|'))).toHaveLength(3)
  })

  it('round-trips the QR-014 table without splitting its response rows', async () => {
    const source = [
      '| Scenario element | Requirement |',
      '| --- | --- |',
      '| Response | The system shall issue target-browser access only after the source browser approves the exact matching code. It shall accept a transfer at most once and only during its five-minute validity period. It shall reject expired, cancelled, reused, code-mismatched, unapproved, and unauthorized transfers, and reject use of a revoked Granted EVCC. Anonymous initiation shall require a valid [Event Visitor Control Credential](../../../terminology/glossary.md#event-visitor-control-credential-evcc), plus an [Event Owner Edit Token](../../../terminology/glossary.md#event-owner-edit-token) when owner authority is requested. |',
      "| Response measure | Automated route and browser tests demonstrate a successful approved transfer and rejection for every invalid transfer state, including source revocation. MongoDB transfer behavior is outside this requirement's scope. |",
      '',
    ].join('\n')

    const once = await format(source)
    const twice = await format(once)

    expect(twice).toBe(once)
    expect(once.split('\n').filter((line) => line.startsWith('|'))).toHaveLength(4)
    expect(once).toContain('five-minute validity period. It shall reject expired')
    expect(once).toContain('including source revocation. MongoDB transfer behavior')
  })

  it('passes custom abbreviations through to the upstream formatter', async () => {
    await expect(
      format('The memo cites Dr. Wu. It records the decision.\n', {
        sentencesPerLineAdditionalAbbreviations: ['Dr.'],
      }),
    ).resolves.toBe(
      'The memo cites Dr. Wu.\nIt records the decision.\n',
    )
  })
})
