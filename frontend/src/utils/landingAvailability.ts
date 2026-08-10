import { isRichLandingEnabled } from "./featureAvailability"

export type { LandingAvailabilityEnvironment } from "./featureAvailability"
export { isRichLandingEnabled } from "./featureAvailability"

export const richLandingEnabled = isRichLandingEnabled(import.meta.env)
