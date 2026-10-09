package processqueue

import (
	"tcp-chat/internal/jobs"
)

type ProcessChannel struct{
	ChannelName string
	Channel chan *jobs.Job

}

func NewProcessChannel (name string) *ProcessChannel{
	return &ProcessChannel{
		ChannelName: name,
		Channel: make(chan *jobs.Job),
	}
}

func (p *ProcessChannel) Push(job *jobs.Job) {
	p.Channel <- job
}


func (p *ProcessChannel) Pop() *jobs.Job{
	job := <- p.Channel
	return job
}