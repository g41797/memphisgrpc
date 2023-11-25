package server

import (
	"context"

	"githib.com/g41797/grpcadapter/pb"
	"github.com/memphisdev/memphis.go"
)

type manageRequest func(req *pb.ManageRequest) (responsible bool, status *pb.Status)

type manager struct {
	bc   *brokerConnector
	proc []manageRequest
	mc   *memphis.Conn
}

func newManager(bc *brokerConnector) *manager {
	manager := new(manager)
	manager.bc = bc
	manager.buildChainOfResp()
	return manager
}

func (srv *manager) Manage(ctx context.Context, mr *pb.ManageRequest) (*pb.Status, error) {
	status := pb.Status{}

	mc, err := srv.bc.connect()
	if err != nil {
		*status.Text = err.Error()
		return &status, err
	}

	srv.mc = mc

	for _, prc := range srv.proc {
		resp, status := prc(mr)
		if resp {
			if status == nil {
				status = &pb.Status{}
			}
			return status, nil
		}
	}

	*status.Text = "not implemented"

	return &status, nil
}

func (srv *manager) buildChainOfResp() {
	if srv.proc != nil {
		return
	}
	srv.proc = append(srv.proc, srv.createStation)
	srv.proc = append(srv.proc, srv.destroyStation)
	srv.proc = append(srv.proc, srv.createProducer)
	srv.proc = append(srv.proc, srv.destroyProducer)
	srv.proc = append(srv.proc, srv.createConsumer)
	srv.proc = append(srv.proc, srv.destroyConsumer)
}

func (srv *manager) createStation(req *pb.ManageRequest) (responsible bool, status *pb.Status) {

	cmd := req.GetCreatestation()
	if cmd != nil {
		_, status := connCreateStation(srv.mc, cmd)
		return true, status
	}

	return false, nil
}

func (srv *manager) destroyStation(req *pb.ManageRequest) (responsible bool, status *pb.Status) {

	cmd := req.GetDestroystation()
	if cmd != nil {
		return true, connDestroyStation(srv.mc, cmd)
	}

	return false, nil
}

func (srv *manager) createProducer(req *pb.ManageRequest) (responsible bool, status *pb.Status) {

	cmd := req.GetCreateproducer()
	if cmd != nil {
		_, status := connCreateProducer(srv.mc, cmd)
		return true, status
	}

	return false, nil
}

func (srv *manager) destroyProducer(req *pb.ManageRequest) (responsible bool, status *pb.Status) {

	cmd := req.GetDestroyproducer()
	if cmd != nil {
		return true, connDestroyProducer(srv.mc, cmd)
	}

	return false, nil
}

func (srv *manager) createConsumer(req *pb.ManageRequest) (responsible bool, status *pb.Status) {

	cmd := req.GetCreateconsumer()
	if cmd != nil {
		_, status := connCreateConsumer(srv.mc, cmd)
		return true, status
	}

	return false, nil
}

func (srv *manager) destroyConsumer(req *pb.ManageRequest) (responsible bool, status *pb.Status) {

	cmd := req.GetDestoyconsumer()
	if cmd != nil {
		return true, connDestroyConsumer(srv.mc, cmd)
	}

	return false, nil
}

func (srv *manager) clean() {
	if srv == nil {
		return
	}

	if srv.mc != nil {
		srv.mc.Close()
	}
}

func connCreateStation(conn *memphis.Conn, req *pb.CreateStationRequest) (string, *pb.Status) {
	return "", nil
}

func connDestroyStation(conn *memphis.Conn, req *pb.DestroyStationRequest) *pb.Status {
	return nil
}

func connCreateProducer(conn *memphis.Conn, req *pb.CreateProducerRequest) (string, *pb.Status) {
	return "", nil
}

func connDestroyProducer(conn *memphis.Conn, req *pb.DestroyProducerRequest) *pb.Status {
	return nil
}

func connCreateConsumer(conn *memphis.Conn, req *pb.CreateConsumerRequest) (string, *pb.Status) {
	return "", nil
}

func connDestroyConsumer(conn *memphis.Conn, req *pb.DestroyConsumerRequest) *pb.Status {
	return nil
}
