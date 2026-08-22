# Quality Requirement Authoring

Read `../README.md` for the shared requirement format, metadata, component,
terminology, and index conventions before creating or changing a quality
requirement.

## ISO/IEC 25010 Classification

This guide uses terms from the [ISO/IEC 25010:2023 product quality model](https://www.iso.org/obp/ui/en/#iso:std:iso-iec:25010:ed-2:v1:en).

Every quality requirement shall declare exactly one ISO/IEC 25010:2023 product
quality characteristic and one of its subcharacteristics in its YAML front
matter:

```yaml
characteristic: security
subcharacteristic: confidentiality
```

Use lowercase values with spaces where the standard uses spaces. The allowed
characteristics and subcharacteristics are:

| Characteristic | Subcharacteristics |
| --- | --- |
| `functional suitability` | `functional completeness`, `functional correctness`, `functional appropriateness` |
| `performance efficiency` | `time behaviour`, `resource utilization`, `capacity` |
| `compatibility` | `co-existence`, `interoperability` |
| `interaction capability` | `appropriateness recognizability`, `learnability`, `operability`, `user error protection`, `user engagement`, `inclusivity`, `user assistance`, `self-descriptiveness` |
| `reliability` | `faultlessness`, `availability`, `fault tolerance`, `recoverability` |
| `security` | `confidentiality`, `integrity`, `non-repudiation`, `accountability`, `authenticity`, `resistance` |
| `maintainability` | `modularity`, `reusability`, `analysability`, `modifiability`, `testability` |
| `flexibility` | `adaptability`, `scalability`, `installability`, `replaceability` |
| `safety` | `operational constraint`, `risk identification`, `fail safe`, `hazard warning`, `safe integration` |

If an outcome addresses more than one subcharacteristic, create separate QR
records. A requirement may refer to related records, but its own front matter
shall retain one classification pair.

## Quality Attribute Scenarios

A QR specifies measurable behavior under stated conditions, rather than a
design tactic or an aspiration such as "fast" or "secure". Its body shall
identify these scenario elements, in prose or a table:

- Source: the user, operator, component, or other actor that causes the event.
- Stimulus: the event or condition that occurs.
- Environment: the relevant load, deployment, event kind, permission state, or
  other condition.
- Artifact: the system component or capability affected.
- Response: the required observable outcome.
- Response measure: the threshold, conformance level, allowed result, or other
  objective criterion that verifies the response.

State a verification approach where the response measure alone does not make
the evidence clear. Use a scenario-specific test, inspection, or operational
exercise; do not prescribe an implementation mechanism unless it is essential
to the quality outcome.

## Applicability And Boundaries

- Name the actor and authority level when the scenario concerns access or
  modification.
- Name the event kind, operation, workload, and fixture size for performance
  scenarios.
- Name the environment for deployment and operational scenarios.
- State exclusions when ambiguity is likely, including development or test
  environments that intentionally differ from staging or production.
- Keep each QR independently understandable and verifiable without its source
  backlog task or an implementation document.

## Review

- Does the front matter contain one valid ISO/IEC 25010 classification pair?
- Does the body provide a measurable scenario with all six elements?
- Can a reviewer identify the verification evidence and relevant boundaries?
- Is the requirement an observable quality outcome rather than a task, design
  decision, or implementation plan?
