package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	requestedFile, fileExists := cb[file]

	occupiedSquareCount := 0

	if fileExists {
		for _, square := range requestedFile {
			if square {
				occupiedSquareCount++
			}
		}
	}

	return occupiedSquareCount
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	occupiedSquareCount := 0

	if (rank < 1 || rank > 8) {
		return occupiedSquareCount
	}

	for _, file := range cb {
		square := file[rank - 1]

		if square {
			occupiedSquareCount++
		}
	}

	return occupiedSquareCount
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	squareCount := 0

	for _, file := range cb {
		squareCount += len(file)
	}

	return squareCount
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupiedSquareCount := 0

	for fileName, _ := range cb {
		occupiedSquareCount += CountInFile(cb, fileName)
	}

	return occupiedSquareCount
}
