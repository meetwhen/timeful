export interface SignInAvailabilityEnvironment {
  VITE_ENABLE_SIGN_IN?: string
}

export interface LandingAvailabilityEnvironment {
  VITE_ENABLE_RICH_LANDING?: string
}

export interface LandingSignInAvailabilityEnvironment
  extends SignInAvailabilityEnvironment, LandingAvailabilityEnvironment {}

function isEnabled(rawValue: string | undefined): boolean {
  const value = rawValue?.trim().toLowerCase()

  if (!value) {
    return true
  }

  return value !== "false"
}

export function isSignInEnabled(
  env: SignInAvailabilityEnvironment = {},
): boolean {
  return isEnabled(env.VITE_ENABLE_SIGN_IN)
}

export function isRichLandingEnabled(
  env: LandingAvailabilityEnvironment = {},
): boolean {
  return isEnabled(env.VITE_ENABLE_RICH_LANDING)
}

export function isLandingSignInEnabled(
  env: LandingSignInAvailabilityEnvironment = {},
): boolean {
  return isSignInEnabled(env) && isRichLandingEnabled(env)
}
