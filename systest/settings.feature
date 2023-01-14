Feature: create a new setting
    As a real human bean,
    I want to add settings to the settingsStore,
    So that I didn't waste a whole weekend writing something that doesn't work

  Scenario: Add a global setting
    Given there are 0 settings in the database
    When I send the default global setting request
    Then there should be 1 global setting in the database
