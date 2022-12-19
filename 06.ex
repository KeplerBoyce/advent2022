contents = File.read!("./inputs/06.txt")
chars = String.split(contents, "", trim: true)

windows = Enum.chunk_every(chars, 4, 1)
try do
  for {w, i} <- Enum.with_index(windows) do
    if Enum.uniq(w) == w do
      IO.puts(i + 4)
      throw(:break)
    end
  end
catch
  :break -> :broken
end

windows = Enum.chunk_every(chars, 14, 1)
try do
  for {w, i} <- Enum.with_index(windows) do
    if Enum.uniq(w) == w do
      IO.puts(i + 14)
      throw(:break)
    end
  end
catch
  :break -> :broken
end
