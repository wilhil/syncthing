// +build !windows,!darwin

package protocol

// Normal Unixes uses NFC and slashes, which is the wire format.

type nativeModel struct {
	next Model
}

func (m nativeModel) Index(nodeID string, repo string, files []FileInfo) {
	m.next.Index(nodeID, repo, files)
}

func (m nativeModel) IndexUpdate(nodeID string, repo string, files []FileInfo) {
	m.next.IndexUpdate(nodeID, repo, files)
}

func (m nativeModel) Request(nodeID, repo string, name string, offset int64, size int) ([]byte, error) {
	return m.next.Request(nodeID, repo, name, offset, size)
}

func (m nativeModel) Close(nodeID string, err error) {
	m.next.Close(nodeID, err)
}
