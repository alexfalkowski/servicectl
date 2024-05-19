Feature: Net

  The ability to verify net.

  Scenario: Sucessfully setup http
    When we try to setup http
    Then we should have a successfully started http
    And I should see a log entry of "http: successfully started" in the file "reports/http.log"

  Scenario: Sucessfully setup grpc
    When we try to setup grpc
    Then we should have a successfully started grpc
    And I should see a log entry of "grpc: successfully started" in the file "reports/grpc.log"
