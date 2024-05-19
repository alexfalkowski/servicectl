Feature: Feature

  The ability to verify featue flags.

  Scenario: Sucessfully connect to feature
    When we try to connect to feature
    Then we should have a succesful feature connection
    And I should see a log entry of "feature: successfully verified connection" in the file "reports/feature.log"
