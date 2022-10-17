package nodes

type Node interface {
	Run(ctx WorkflowContext, input any) (any, error)
}
