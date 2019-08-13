package runcontainerd

import (
	"io"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/rundmc/goci"
	"code.cloudfoundry.org/lager"
)

//go:generate counterfeiter . PeaManager
type PeaManager interface {
	Create(log lager.Logger, id string, bundle goci.Bndl, io garden.ProcessIO) error
	Delete(log lager.Logger, containerID string) error
	RemoveBundle(log lager.Logger, containerID string) error
}

//go:generate counterfeiter . Volumizer
type Volumizer interface {
	Destroy(log lager.Logger, handle string) error
}

type RunContainerPea struct {
	PeaManager     PeaManager
	ProcessManager ProcessManager
	Volumizer      Volumizer
}

func (r *RunContainerPea) RunPea(
	log lager.Logger, processID string, processBundle goci.Bndl, sandboxHandle string,
	pio garden.ProcessIO, tty bool, procJSON io.Reader, extraCleanup func() error,
) (garden.Process, error) {

	if processBundle.Spec.Annotations == nil {
		processBundle.Spec.Annotations = make(map[string]string)
	}
	processBundle.Spec.Annotations["container-type"] = "pea"
	processBundle.Spec.Annotations["sandbox-container"] = sandboxHandle

	log.Info("using runcontainerd.RunPea")

	if err := r.PeaManager.Create(log, processID, processBundle, pio); err != nil {
		return &Process{}, err
	}

	// TODO fix
	process, err := r.ProcessManager.GetTask(log, map[string]string{"container-type": "pea", "sandbox-container": sandboxHandle}, processID)
	if err != nil {
		return nil, err
	}

	return NewPeaProcess(log, process, r.PeaManager, r.Volumizer), nil
}
