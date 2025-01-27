Feature: Cache
  The ability to verify database.

  Scenario: Successfully connect to pg
    When we try to connect to pg
    Then we should have a succesful pg connection
    And I should see a log entry of "pg: successfully verified connection" in the file "reports/pg.log"
