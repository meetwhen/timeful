# CAND-139

#### Source

> - [x] Given I sign in, when I enter an unregistered email and click Continue with email:
>   - The input field is highlighted red
>   - The error appears like on accounts.google.com: (red alert icon) "Couldn’t find this account. Create account"

[Source lines 479-481](../../../../backlog/backlog.md#L479-L481)

#### Candidate behavior

An unregistered email submission highlights the input red and shows `Couldn’t find this account. Create account` with a red alert icon.

#### Applicability

Actor: sign-in visitor; Location: Sign in page email flow; Event kind: not applicable; Interaction mode: submit unregistered email; Viewport: unspecified; State: email has no account; Exclusions: registered email and other providers.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-023 displays email registration status during sign-in; this source supplies its unregistered-email presentation. Confidence: inferred.

#### Disposition

Review as a presentation refinement of FR-023.

#### Open Questions

Is the exact accounts.google.com-like visual treatment a durable requirement, and does Create account initiate registration?
