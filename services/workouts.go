package services

type body_part int64

const (
	chest      body_part = 0
	back                 = 1
	quads                = 2
	hamstrings           = 3
	glutes               = 4
	biceps               = 5
	triceps              = 6
	cardio               = 7
)

type exercise struct {
	name   string
	part   body_part
	reps   int64
	weight int64
	order  int64
	link   string
}

type movement struct {
	exercises []exercise
	cycles    int64
	order     int64
}

type workout struct {
	name      string
	movements []movement
	notes     string
	order     int64
}

// WorkoutRoutine is a workout plan that is meant
type WorkoutRoutine struct {
	Name     string
	workouts []workout
}

func GetRoutines() []*WorkoutRoutine {
	var ws [](*WorkoutRoutine)
	return ws
}

func CreateRoutine() *WorkoutRoutine {
	return new(WorkoutRoutine)
}
