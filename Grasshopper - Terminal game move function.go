package kata

func Move(position int, roll int) int {
	position += roll * 2
	return position
}
