package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	fl, ok := cb[file]
	if !ok {
		return 0
	}
	c := 0
	for iv := range fl {
		if fl[iv] {
			c += 1
		}
	}
	return c
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	c := 0
	if rank < 1 || rank > 8 {
		return 0
	}
	for _, fl := range cb {
		if fl[rank - 1] {
			c += 1
		}
	}
	return c
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	cs := 0
	for _,fl := range cb {
		for range fl {
			cs += 1
		}
	}
	return cs
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	cs := 0
	for _,fl := range cb {
		for _,item := range fl {
			if item {
				cs += 1
			}
		}
	}
	return cs
}
