Feature: Cache

  The ability to verify cache.

  Scenario: Sucessfully connect to redis
    When we try to connect to redis
    Then we should have a succesful redis connection
    And I should see a log entry of "redis: successfully verified connection" in the file "reports/redis.log"
