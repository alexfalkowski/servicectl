Feature: Security

  The ability to manage security keys.

  Scenario: Sucessfully verify a key
    When we try to verify a key
    Then we should have a succesful key verification
    And I should see a log entry of "token: successfully verified key" in the file "reports/security.log"

  Scenario: Sucessfully rotate a key
    When we try to rotate a key
    Then we should have a succesfully rotated the key
    And I should see a log entry of "token: successfully rotated key and hash" in the file "reports/security.log"
