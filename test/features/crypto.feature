Feature: Crypto
  The ability to rotate keys.

  Scenario: Successfully rotate AES key
    When we rotate the aes key
    Then we should have a succesful rotation of aes keys
    And I should see a log entry of "aes: successfully rotated key" in the file "reports/aes.log"

  Scenario: Successfully verify AES key
    When we verify the aes key
    Then we should have a succesful verification of aes keys
    And I should see a log entry of "aes: successfully verified key" in the file "reports/aes.log"

  Scenario: Successfully rotate HMAC key
    When we rotate the hmac key
    Then we should have a succesful rotation of hmac keys
    And I should see a log entry of "hmac: successfully rotated key" in the file "reports/hmac.log"

  Scenario: Successfully verify HMAC key
    When we verify the hmac key
    Then we should have a succesful verification of hmac keys
    And I should see a log entry of "hmac: successfully verified key" in the file "reports/hmac.log"

  Scenario: Successfully rotate RSA key
    When we rotate the rsa key
    Then we should have a succesful rotation of rsa keys
    And I should see a log entry of "rsa: successfully rotated keys" in the file "reports/rsa.log"

  Scenario: Successfully verify RSA key
    When we verify the rsa key
    Then we should have a succesful verification of rsa keys
    And I should see a log entry of "rsa: successfully verified keys" in the file "reports/rsa.log"

  Scenario: Successfully rotate Ed25519 key
    When we rotate the ed25519 key
    Then we should have a succesful rotation of ed25519 keys
    And I should see a log entry of "ed25519: successfully rotated keys" in the file "reports/ed25519.log"

  Scenario: Successfully verify Ed25519 key
    When we verify the ed25519 key
    Then we should have a succesful verification of ed25519 keys
    And I should see a log entry of "ed25519: successfully verified keys" in the file "reports/ed25519.log"
