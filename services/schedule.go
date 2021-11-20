package services




type ScheuleWorkoutRoutine struct{
	CurrentRoutine WorkoutRoutine
	NextWorkout int64 // order number of the next workout
}



type Schedule struct{
	ScheuleWorkoutRoutine

}