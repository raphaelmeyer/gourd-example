
Feature: Enter numbers
  As a ...
  I want to ...
  In order to ...

#  _______
#  _______
#
#  7 8 9 /
#  4 5 6 *
#  1 2 3 -
#  0 C = +
#

Scenario: Cleared screen
  Given the calculator is cleared
  Then the result on the screen should be 0

Scenario: Press zero when calculator shows zero
  Given the calculator is cleared
  When I press 0
  Then the result on the screen should be 0

Scenario: Press a number when calculator shows zero
  Given the calculator is cleared
  When I press 1
  Then the result on the screen should be 1

Scenario: Enter three digit number
  Given the calculator is cleared
  When I press 1
  And I press 3
  And I press 5
  Then the result on the screen should be 135

