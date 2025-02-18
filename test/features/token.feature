Feature: Token
  The ability to rotate token secrets.

  Scenario Outline: Successfully rotate tokens
    When we rotate token secret of kind "<kind>"
    Then we should have a succesful rotation of the secret of kind "<kind>"
    And I should see a log entry of "token: successfully rotated <kind>" in the file "reports/token.log"

    Examples:
      | kind   |
      | opaque |
