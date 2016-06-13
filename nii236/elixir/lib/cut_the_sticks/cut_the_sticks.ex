defmodule CutTheSticks do
  def main() do
    a = IO.gets
    b = IO.gets
    process("#{a}\n#{b}")
  end
  
  def process(input) do
    values = String.split(input, "\n")
      |> List.delete_at(0)
      |> List.delete_at(-1)
      |> List.to_string
      |> String.strip
      |> String.split(" ")
      |> Enum.map(&(Integer.parse(&1)))
      |> Enum.map(&(elem(&1,0)))
    recurse_cut(values, Enum.max(values))
  end

  def recurse_cut(values, max) when max > 0 do
    min = values
      |> Enum.filter(&(&1 > 0))
      |> Enum.min
    IO.puts(Enum.count(values, &(&1 >= min)))
    result = Enum.map(values, &(&1-min))
    recurse_cut(result, Enum.max(result))
  end

  def recurse_cut(_values, max) when max <= 0 do
  end
end
