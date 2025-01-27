Feature: Hooks
  The ability to verify hooks.

  Scenario: Successfully verify a hook
    When we try to verify a hook
    Then we should have a succesful hook verification
    And I should see a log entry of "hooks: successfully verified" in the file "reports/hooks.log"

  Scenario: Successfully rotate a hook secret
    When we try to rotate a secret for the hook
    Then we should have a succesful rotated the hook secret
    And I should see a log entry of "hooks: successfully rotated" in the file "reports/hooks.log"
