import { isSignInEnabled } from "./featureAvailability"

export type { SignInAvailabilityEnvironment } from "./featureAvailability"
export { isSignInEnabled } from "./featureAvailability"

export const signInEnabled = isSignInEnabled(import.meta.env)
