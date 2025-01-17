package executables_test

import (
	"context"

	"github.com/golang/mock/gomock"

	"github.com/aws/eks-anywhere/pkg/executables"
	"github.com/aws/eks-anywhere/pkg/executables/mocks"
)

type commandExpect struct {
	command *executables.Command
	e       *mocks.MockExecutable
}

func expectCommand(e *mocks.MockExecutable, ctx context.Context, args ...string) *commandExpect {
	e.EXPECT().Command(ctx, args).Return(executables.NewCommand(ctx, e, args...))
	return &commandExpect{
		command: executables.NewCommand(ctx, e, args...),
		e:       e,
	}
}

func (c *commandExpect) withEnvVars(envVars map[string]string) *commandExpect {
	c.command.WithEnvVars(envVars)
	return c
}

func (c *commandExpect) withStdIn(stdIn []byte) *commandExpect {
	c.command.WithStdIn(stdIn)
	return c
}

func (c *commandExpect) to() *gomock.Call {
	return c.e.EXPECT().Run(c.command)
}
