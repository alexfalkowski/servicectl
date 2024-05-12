Feature: Crypto

  The ability to rotate keys.

  Scenario: Succesfuly rotate AES key
    When we rotate the aes key
    Then we should have a succesful rotation of aes keys
    And I should see a log entry of "aes: successfully rotated key" in the file "reports/aes.log"

  Scenario: Succesfuly verify AES key
    When we verify the aes key
    Then I should see a log entry of "aes: successfully verified key" in the file "reports/aes.log"

  Scenario: Succesfuly rotate HMAC key
    When we rotate the hmac key
    Then we should have a succesful rotation of hmac keys
    And I should see a log entry of "hmac: successfully rotated key" in the file "reports/hmac.log"

  Scenario: Succesfuly verify HMAC key
    When we verify the hmac key
    Then I should see a log entry of "hmac: successfully verified key" in the file "reports/hmac.log"
