package protoprocess

import (
	"fmt"
	"io"

	"go.pedge.io/pb/go/google/protobuf"
	"go.pedge.io/proto/stream"
)

func send(reader io.Reader, streamingBytesServer protostream.StreamingBytesServer, chunkSizeBytes int) error {
	p := make([]byte, chunkSizeBytes)
	count := 0
	for n, err := reader.Read(p); err != io.EOF; n, err = reader.Read(p) {
		if err != nil {
			return err
		}
		if n < chunkSizeBytes {
			p = p[:n]
		}
		if err := streamingBytesServer.Send(
			&google_protobuf.BytesValue{
				Value: p,
			},
		); err != nil {
			return err
		}
		count += n
		p = make([]byte, chunkSizeBytes)
	}
	return nil
}

func recv(writer io.Writer, streamingBytesClient protostream.StreamingBytesClient, chunkSizeBytes int) error {
	count := 0
	for bytesValue, err := streamingBytesClient.Recv(); err != io.EOF; bytesValue, err = streamingBytesClient.Recv() {
		if err != nil {
			return err
		}
		if bytesValue != nil && bytesValue.Value != nil && len(bytesValue.Value) > 0 {
			count += len(bytesValue.Value)
			n, err := writer.Write(bytesValue.Value)
			if err != nil {
				return err
			}
			if n != len(bytesValue.Value) {
				return fmt.Errorf("protoprocess: tried to write %d bytes, wrote %d", len(bytesValue.Value), n)
			}
		}
	}
	return nil
}
