import System.IO
import Control.Monad

main = do
    file <- readFile $ "./inputs/02.txt"
    let rows = lines file
    print . sum $ map (rpc . halfSplit) rows
    print . sum $ map (rpc2 . halfSplit) rows

rpc :: (String, String) -> Integer
rpc (a, b)
    | a == "B" && b == "X" = 1
    | a == "A" && b == "Z" = 3
    | a == "A" && b == "X" = 4
    | a == "B" && b == "Y" = 5
    | a == "C" && b == "Y" = 2
    | a == "C" && b == "Z" = 6
    | a == "C" && b == "X" = 7
    | a == "A" && b == "Y" = 8
    | a == "B" && b == "Z" = 9

rpc2 :: (String, String) -> Integer
rpc2 (a, b)
    | a == "B" && b == "X" = 1
    | a == "C" && b == "X" = 2
    | a == "A" && b == "X" = 3
    | a == "A" && b == "Y" = 4
    | a == "B" && b == "Y" = 5
    | a == "C" && b == "Y" = 6
    | a == "C" && b == "Z" = 7
    | a == "A" && b == "Z" = 8
    | a == "B" && b == "Z" = 9

halfSplit :: String -> (String, String)
halfSplit str = (head . words $ str, last . words $ str)
