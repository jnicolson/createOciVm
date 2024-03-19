package main

import (
	"context"

	"github.com/go-co-op/gocron/v2"
	"github.com/rs/zerolog/log"
)

type Schedule struct {
	sc       gocron.Scheduler
	schedule string
	job      gocron.Job
}

func NewSchedule(ctx context.Context, schedule string) *Schedule {
	return &Schedule{
		schedule: schedule,
	}
}

func (s *Schedule) Start() {
	var err error
	s.sc, err = gocron.NewScheduler(gocron.WithLogger(NewJLogger()))
	if err != nil {
		log.Fatal().Err(err).Msg("Creating scheduler")
	}

	s.sc.Start()

}

func (s *Schedule) Stop() {
	s.sc.Shutdown()
}

func (s *Schedule) CancelJob() {
	s.sc.RemoveJob(s.job.ID())
	log.Info().Str("ID", s.job.ID().String()).Msg("Job Stopped")
}

func (s *Schedule) SetJob(job func()) {
	cronJob := gocron.CronJob(s.schedule, true)
	//task := gocron.NewTask(func() { log.Info().Msg("in gocron") })
	task := gocron.NewTask(job)
	j, err := s.sc.NewJob(cronJob, task)

	if err != nil {
		log.Fatal().Err(err).Msg("Error creating job")
	}

	s.job = j
	log.Info().Str("ID", s.job.ID().String()).Msg("Job started")

}
