Feature: Hooks

  The ability to verify hooks.

  Scenario: Sucessfully sign a hook
    When we try to sign a hook
    Then we should have a succesful hook signature
    And I should see a log entry of "hooks: successfully signed" in the file "reports/hooks.log"
