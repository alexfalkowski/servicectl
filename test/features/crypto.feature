Feature: Crypto

  The ability to rotate keys.

  Scenario: Succesfuly rotate AES key
    When we rotate the aes key
    Then it should run sucessfully
    And we should have a config file "reports/aes_client.yml"
    And I should see a log entry of "rotated aes key" in the file "reports/aes.log"
