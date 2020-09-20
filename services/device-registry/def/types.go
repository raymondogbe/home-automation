// Code generated by jrpc. DO NOT EDIT.

package deviceregistrydef

// DeviceHeader is defined in the .def file
type DeviceHeader struct {
	Id             *string                 `json:"id,omitempty"`
	Name           *string                 `json:"name,omitempty"`
	Type           *string                 `json:"type,omitempty"`
	Kind           *string                 `json:"kind,omitempty"`
	ControllerName *string                 `json:"controller_name,omitempty"`
	Attributes     *map[string]interface{} `json:"attributes,omitempty"`
	StateProviders *[]string               `json:"state_providers,omitempty"`
	RoomId         *string                 `json:"room_id,omitempty"`
	Room           *Room                   `json:"room,omitempty"`
}

// GetId returns the de-referenced value of Id.
// The second return value states whether the field was set.
func (m *DeviceHeader) GetId() (val string, set bool) {
	if m.Id == nil {
		return
	}

	return *m.Id, true
}

// SetId sets the value of Id
func (m *DeviceHeader) SetId(v string) *DeviceHeader {
	m.Id = &v
	return m
}

// GetName returns the de-referenced value of Name.
// The second return value states whether the field was set.
func (m *DeviceHeader) GetName() (val string, set bool) {
	if m.Name == nil {
		return
	}

	return *m.Name, true
}

// SetName sets the value of Name
func (m *DeviceHeader) SetName(v string) *DeviceHeader {
	m.Name = &v
	return m
}

// GetType returns the de-referenced value of Type.
// The second return value states whether the field was set.
func (m *DeviceHeader) GetType() (val string, set bool) {
	if m.Type == nil {
		return
	}

	return *m.Type, true
}

// SetType sets the value of Type
func (m *DeviceHeader) SetType(v string) *DeviceHeader {
	m.Type = &v
	return m
}

// GetKind returns the de-referenced value of Kind.
// The second return value states whether the field was set.
func (m *DeviceHeader) GetKind() (val string, set bool) {
	if m.Kind == nil {
		return
	}

	return *m.Kind, true
}

// SetKind sets the value of Kind
func (m *DeviceHeader) SetKind(v string) *DeviceHeader {
	m.Kind = &v
	return m
}

// GetControllerName returns the de-referenced value of ControllerName.
// The second return value states whether the field was set.
func (m *DeviceHeader) GetControllerName() (val string, set bool) {
	if m.ControllerName == nil {
		return
	}

	return *m.ControllerName, true
}

// SetControllerName sets the value of ControllerName
func (m *DeviceHeader) SetControllerName(v string) *DeviceHeader {
	m.ControllerName = &v
	return m
}

// GetAttributes returns the de-referenced value of Attributes.
// The second return value states whether the field was set.
func (m *DeviceHeader) GetAttributes() (val map[string]interface{}, set bool) {
	if m.Attributes == nil {
		return
	}

	return *m.Attributes, true
}

// SetAttributes sets the value of Attributes
func (m *DeviceHeader) SetAttributes(v map[string]interface{}) *DeviceHeader {
	m.Attributes = &v
	return m
}

// GetStateProviders returns the de-referenced value of StateProviders.
// The second return value states whether the field was set.
func (m *DeviceHeader) GetStateProviders() (val []string, set bool) {
	if m.StateProviders == nil {
		return
	}

	return *m.StateProviders, true
}

// SetStateProviders sets the value of StateProviders
func (m *DeviceHeader) SetStateProviders(v []string) *DeviceHeader {
	m.StateProviders = &v
	return m
}

// GetRoomId returns the de-referenced value of RoomId.
// The second return value states whether the field was set.
func (m *DeviceHeader) GetRoomId() (val string, set bool) {
	if m.RoomId == nil {
		return
	}

	return *m.RoomId, true
}

// SetRoomId sets the value of RoomId
func (m *DeviceHeader) SetRoomId(v string) *DeviceHeader {
	m.RoomId = &v
	return m
}

// GetRoom returns the de-referenced value of Room.
// The second return value states whether the field was set.
func (m *DeviceHeader) GetRoom() (val Room, set bool) {
	if m.Room == nil {
		return
	}

	return *m.Room, true
}

// SetRoom sets the value of Room
func (m *DeviceHeader) SetRoom(v Room) *DeviceHeader {
	m.Room = &v
	return m
}

// Validate returns an error if any of the fields have bad values
func (m *DeviceHeader) Validate() error {
	if err := m.Room.Validate(); err != nil {
		return err
	}

	return nil
}

// Room is defined in the .def file
type Room struct {
	Id      *string         `json:"id,omitempty"`
	Name    *string         `json:"name,omitempty"`
	Devices *[]DeviceHeader `json:"devices,omitempty"`
}

// GetId returns the de-referenced value of Id.
// The second return value states whether the field was set.
func (m *Room) GetId() (val string, set bool) {
	if m.Id == nil {
		return
	}

	return *m.Id, true
}

// SetId sets the value of Id
func (m *Room) SetId(v string) *Room {
	m.Id = &v
	return m
}

// GetName returns the de-referenced value of Name.
// The second return value states whether the field was set.
func (m *Room) GetName() (val string, set bool) {
	if m.Name == nil {
		return
	}

	return *m.Name, true
}

// SetName sets the value of Name
func (m *Room) SetName(v string) *Room {
	m.Name = &v
	return m
}

// GetDevices returns the de-referenced value of Devices.
// The second return value states whether the field was set.
func (m *Room) GetDevices() (val []DeviceHeader, set bool) {
	if m.Devices == nil {
		return
	}

	return *m.Devices, true
}

// SetDevices sets the value of Devices
func (m *Room) SetDevices(v []DeviceHeader) *Room {
	m.Devices = &v
	return m
}

// Validate returns an error if any of the fields have bad values
func (m *Room) Validate() error {
	if m.Devices != nil {
		for _, r := range *m.Devices {
			if err := r.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetDeviceRequest is defined in the .def file
type GetDeviceRequest struct {
	DeviceId *string `json:"device_id,omitempty"`
}

// GetDeviceId returns the de-referenced value of DeviceId.
// The second return value states whether the field was set.
func (m *GetDeviceRequest) GetDeviceId() (val string, set bool) {
	if m.DeviceId == nil {
		return
	}

	return *m.DeviceId, true
}

// SetDeviceId sets the value of DeviceId
func (m *GetDeviceRequest) SetDeviceId(v string) *GetDeviceRequest {
	m.DeviceId = &v
	return m
}

// Validate returns an error if any of the fields have bad values
func (m *GetDeviceRequest) Validate() error {
	return nil
}

// GetDeviceResponse is defined in the .def file
type GetDeviceResponse struct {
	DeviceHeader *DeviceHeader `json:"device_header,omitempty"`
}

// GetDeviceHeader returns the de-referenced value of DeviceHeader.
// The second return value states whether the field was set.
func (m *GetDeviceResponse) GetDeviceHeader() (val DeviceHeader, set bool) {
	if m.DeviceHeader == nil {
		return
	}

	return *m.DeviceHeader, true
}

// SetDeviceHeader sets the value of DeviceHeader
func (m *GetDeviceResponse) SetDeviceHeader(v DeviceHeader) *GetDeviceResponse {
	m.DeviceHeader = &v
	return m
}

// Validate returns an error if any of the fields have bad values
func (m *GetDeviceResponse) Validate() error {
	if err := m.DeviceHeader.Validate(); err != nil {
		return err
	}

	return nil
}

// ListDevicesRequest is defined in the .def file
type ListDevicesRequest struct {
	ControllerName *string `json:"controller_name,omitempty"`
}

// GetControllerName returns the de-referenced value of ControllerName.
// The second return value states whether the field was set.
func (m *ListDevicesRequest) GetControllerName() (val string, set bool) {
	if m.ControllerName == nil {
		return
	}

	return *m.ControllerName, true
}

// SetControllerName sets the value of ControllerName
func (m *ListDevicesRequest) SetControllerName(v string) *ListDevicesRequest {
	m.ControllerName = &v
	return m
}

// Validate returns an error if any of the fields have bad values
func (m *ListDevicesRequest) Validate() error {
	return nil
}

// ListDevicesResponse is defined in the .def file
type ListDevicesResponse struct {
	DeviceHeaders *[]DeviceHeader `json:"device_headers,omitempty"`
}

// GetDeviceHeaders returns the de-referenced value of DeviceHeaders.
// The second return value states whether the field was set.
func (m *ListDevicesResponse) GetDeviceHeaders() (val []DeviceHeader, set bool) {
	if m.DeviceHeaders == nil {
		return
	}

	return *m.DeviceHeaders, true
}

// SetDeviceHeaders sets the value of DeviceHeaders
func (m *ListDevicesResponse) SetDeviceHeaders(v []DeviceHeader) *ListDevicesResponse {
	m.DeviceHeaders = &v
	return m
}

// Validate returns an error if any of the fields have bad values
func (m *ListDevicesResponse) Validate() error {
	if m.DeviceHeaders != nil {
		for _, r := range *m.DeviceHeaders {
			if err := r.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetRoomRequest is defined in the .def file
type GetRoomRequest struct {
	RoomId *string `json:"room_id,omitempty"`
}

// GetRoomId returns the de-referenced value of RoomId.
// The second return value states whether the field was set.
func (m *GetRoomRequest) GetRoomId() (val string, set bool) {
	if m.RoomId == nil {
		return
	}

	return *m.RoomId, true
}

// SetRoomId sets the value of RoomId
func (m *GetRoomRequest) SetRoomId(v string) *GetRoomRequest {
	m.RoomId = &v
	return m
}

// Validate returns an error if any of the fields have bad values
func (m *GetRoomRequest) Validate() error {
	return nil
}

// GetRoomResponse is defined in the .def file
type GetRoomResponse struct {
	Room *Room `json:"room,omitempty"`
}

// GetRoom returns the de-referenced value of Room.
// The second return value states whether the field was set.
func (m *GetRoomResponse) GetRoom() (val Room, set bool) {
	if m.Room == nil {
		return
	}

	return *m.Room, true
}

// SetRoom sets the value of Room
func (m *GetRoomResponse) SetRoom(v Room) *GetRoomResponse {
	m.Room = &v
	return m
}

// Validate returns an error if any of the fields have bad values
func (m *GetRoomResponse) Validate() error {
	if err := m.Room.Validate(); err != nil {
		return err
	}

	return nil
}

// ListRoomsRequest is defined in the .def file
type ListRoomsRequest struct {
}

// Validate returns an error if any of the fields have bad values
func (m *ListRoomsRequest) Validate() error {
	return nil
}

// ListRoomsResponse is defined in the .def file
type ListRoomsResponse struct {
	Rooms *[]Room `json:"rooms,omitempty"`
}

// GetRooms returns the de-referenced value of Rooms.
// The second return value states whether the field was set.
func (m *ListRoomsResponse) GetRooms() (val []Room, set bool) {
	if m.Rooms == nil {
		return
	}

	return *m.Rooms, true
}

// SetRooms sets the value of Rooms
func (m *ListRoomsResponse) SetRooms(v []Room) *ListRoomsResponse {
	m.Rooms = &v
	return m
}

// Validate returns an error if any of the fields have bad values
func (m *ListRoomsResponse) Validate() error {
	if m.Rooms != nil {
		for _, r := range *m.Rooms {
			if err := r.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
