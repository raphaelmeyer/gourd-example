
Feature: Addition
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

Scenario: Add two numbers
  Given the calculator is cleared
  When I press 5
  And  I press +
  And I press 7
  And I press =
  Then the result on the screen should be 12


