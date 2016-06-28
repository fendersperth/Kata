require(CutTheSticks)

defmodule CutTheSticksTest do
  use ExUnit.Case, async: true
  doctest CutTheSticks

  test "cases from hacker rank" do
    cases = ["6\n5 4 4 2 2 8\n", "8\n1 2 3 4 3 3 2 1\n"]
    Enum.each(cases, &(CutTheSticks.process(&1)))
  end
end
