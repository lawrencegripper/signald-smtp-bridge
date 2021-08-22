package v1

// DO NOT EDIT: this file is automatically generated by ./tools/generator in this repo

import (
	"encoding/json"
	"fmt"
	"log"

	"gitlab.com/signald/signald-go/signald"
)

// Submit: Accept a v2 group invitation. Note that you must have a profile name set to join groups.
func (r *AcceptInvitationRequest) Submit(conn *signald.Signald) (response JsonGroupV2Info, err error) {
	r.Version = "v1"
	r.Type = "accept_invitation"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: Link a new device to a local Signal account
func (r *AddLinkedDeviceRequest) Submit(conn *signald.Signald) (err error) {
	r.Version = "v1"
	r.Type = "add_device"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	return err

}

// Submit: approve a request to join a group
func (r *ApproveMembershipRequest) Submit(conn *signald.Signald) (response JsonGroupV2Info, err error) {
	r.Version = "v1"
	r.Type = "approve_membership"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

func (r *CreateGroupRequest) Submit(conn *signald.Signald) (response JsonGroupV2Info, err error) {
	r.Version = "v1"
	r.Type = "create_group"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: delete all account data signald has on disk, and optionally delete the account from the server as well. Note that this is not "unlink" and will delete the entire account, even from a linked device.
func (r *DeleteAccountRequest) Submit(conn *signald.Signald) (err error) {
	r.Version = "v1"
	r.Type = "delete_account"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	return err

}

// Submit: After a linking URI has been requested, finish_link must be called with the session_id provided with the URI. it will return information about the new account once the linking process is completed by the other device.
func (r *FinishLinkRequest) Submit(conn *signald.Signald) (response Account, err error) {
	r.Version = "v1"
	r.Type = "finish_link"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: Generate a linking URI. Typically this is QR encoded and scanned by the primary device. Submit the returned session_id with a finish_link request.
func (r *GenerateLinkingURIRequest) Submit(conn *signald.Signald) (response LinkingURI, err error) {
	r.Version = "v1"
	r.Type = "generate_linking_uri"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: get all known identity keys
func (r *GetAllIdentities) Submit(conn *signald.Signald) (response AllIdentityKeyList, err error) {
	r.Version = "v1"
	r.Type = "get_all_identities"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: Query the server for the latest state of a known group. If no account in signald is a member of the group (anymore), an error with error_type: 'UnknownGroupException' is returned.
func (r *GetGroupRequest) Submit(conn *signald.Signald) (response JsonGroupV2Info, err error) {
	r.Version = "v1"
	r.Type = "get_group"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: Get information about a known keys for a particular address
func (r *GetIdentitiesRequest) Submit(conn *signald.Signald) (response IdentityKeyList, err error) {
	r.Version = "v1"
	r.Type = "get_identities"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: list all linked devices on a Signal account
func (r *GetLinkedDevicesRequest) Submit(conn *signald.Signald) (response LinkedDevices, err error) {
	r.Version = "v1"
	r.Type = "get_linked_devices"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: Get all information available about a user
func (r *GetProfileRequest) Submit(conn *signald.Signald) (response Profile, err error) {
	r.Version = "v1"
	r.Type = "get_profile"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: Get information about a group from a signal.group link
func (r *GroupLinkInfoRequest) Submit(conn *signald.Signald) (response JsonGroupJoinInfo, err error) {
	r.Version = "v1"
	r.Type = "group_link_info"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: Join a group using the a signal.group URL. Note that you must have a profile name set to join groups.
func (r *JoinGroupRequest) Submit(conn *signald.Signald) (response JsonGroupJoinInfo, err error) {
	r.Version = "v1"
	r.Type = "join_group"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

func (r *LeaveGroupRequest) Submit(conn *signald.Signald) (response GroupInfo, err error) {
	r.Version = "v1"
	r.Type = "leave_group"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: return all local accounts
func (r *ListAccountsRequest) Submit(conn *signald.Signald) (response AccountList, err error) {
	r.Version = "v1"
	r.Type = "list_accounts"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

func (r *ListContactsRequest) Submit(conn *signald.Signald) (response ProfileList, err error) {
	r.Version = "v1"
	r.Type = "list_contacts"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

func (r *ListGroupsRequest) Submit(conn *signald.Signald) (response GroupList, err error) {
	r.Version = "v1"
	r.Type = "list_groups"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

func (r *MarkReadRequest) Submit(conn *signald.Signald) (err error) {
	r.Version = "v1"
	r.Type = "mark_read"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	return err

}

// Submit: react to a previous message
func (r *ReactRequest) Submit(conn *signald.Signald) (response SendResponse, err error) {
	r.Version = "v1"
	r.Type = "react"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: begin the account registration process by requesting a phone number verification code. when the code is received, submit it with a verify request
func (r *RegisterRequest) Submit(conn *signald.Signald) (response Account, err error) {
	r.Version = "v1"
	r.Type = "register"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: delete a message previously sent
func (r *RemoteDeleteRequest) Submit(conn *signald.Signald) (response SendResponse, err error) {
	r.Version = "v1"
	r.Type = "remote_delete"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: Remove a linked device from the Signal account. Only allowed when the local device id is 1
func (r *RemoveLinkedDeviceRequest) Submit(conn *signald.Signald) (err error) {
	r.Version = "v1"
	r.Type = "remove_linked_device"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	return err

}

// Submit: Request other devices on the account send us their group list, syncable config and contact list.
func (r *RequestSyncRequest) Submit(conn *signald.Signald) (err error) {
	r.Version = "v1"
	r.Type = "request_sync"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	return err

}

// Submit: reset a session with a particular user
func (r *ResetSessionRequest) Submit(conn *signald.Signald) (response SendResponse, err error) {
	r.Version = "v1"
	r.Type = "reset_session"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: Resolve a partial JsonAddress with only a number or UUID to one with both. Anywhere that signald accepts a JsonAddress will except a partial, this is a convenience function for client authors, mostly because signald doesn't resolve all the partials it returns
func (r *ResolveAddressRequest) Submit(conn *signald.Signald) (response JsonAddress, err error) {
	r.Version = "v1"
	r.Type = "resolve_address"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

func (r *SendRequest) Submit(conn *signald.Signald) (response SendResponse, err error) {
	r.Version = "v1"
	r.Type = "send"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: set this device's name. This will show up on the mobile device on the same account under
func (r *SetDeviceNameRequest) Submit(conn *signald.Signald) (err error) {
	r.Version = "v1"
	r.Type = "set_device_name"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	return err

}

// Submit: Set the message expiration timer for a thread. Expiration must be specified in seconds, set to 0 to disable timer
func (r *SetExpirationRequest) Submit(conn *signald.Signald) (response SendResponse, err error) {
	r.Version = "v1"
	r.Type = "set_expiration"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

func (r *SetProfile) Submit(conn *signald.Signald) (err error) {
	r.Version = "v1"
	r.Type = "set_profile"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	return err

}

// Submit: receive incoming messages. After making a subscribe request, incoming messages will be sent to the client encoded as ClientMessageWrapper. Send an unsubscribe request or disconnect from the socket to stop receiving messages.
func (r *SubscribeRequest) Submit(conn *signald.Signald) (err error) {
	r.Version = "v1"
	r.Type = "subscribe"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	return err

}

// Submit: Trust another user's safety number using either the QR code data or the safety number text
func (r *TrustRequest) Submit(conn *signald.Signald) (err error) {
	r.Version = "v1"
	r.Type = "trust"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	return err

}

// Submit: send a typing started or stopped message
func (r *TypingRequest) Submit(conn *signald.Signald) (err error) {
	r.Version = "v1"
	r.Type = "typing"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	return err

}

// Submit: See subscribe for more info
func (r *UnsubscribeRequest) Submit(conn *signald.Signald) (err error) {
	r.Version = "v1"
	r.Type = "unsubscribe"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	return err

}

// Submit: update information about a local contact
func (r *UpdateContactRequest) Submit(conn *signald.Signald) (response Profile, err error) {
	r.Version = "v1"
	r.Type = "update_contact"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: modify a group. Note that only one modification action may be preformed at once
func (r *UpdateGroupRequest) Submit(conn *signald.Signald) (response GroupInfo, err error) {
	r.Version = "v1"
	r.Type = "update_group"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

// Submit: verify an account's phone number with a code after registering, completing the account creation process
func (r *VerifyRequest) Submit(conn *signald.Signald) (response Account, err error) {
	r.Version = "v1"
	r.Type = "verify"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}

func (r *VersionRequest) Submit(conn *signald.Signald) (response JsonVersionMessage, err error) {
	r.Version = "v1"
	r.Type = "version"
	if r.ID == "" {
		r.ID = signald.GenerateID()
	}
	err = conn.RawRequest(r)
	if err != nil {
		log.Println("signald-go: error submitting request to signald")
		return
	}

	responseChannel := conn.GetResponseListener(r.ID)
	defer conn.CloseResponseListener(r.ID)

	rawResponse := <-responseChannel
	if rawResponse.Error != nil {
		err = fmt.Errorf("signald error: %s", string(rawResponse.Error))
		return
	}

	err = json.Unmarshal(rawResponse.Data, &response)
	if err != nil {
		rawResponseJson, _ := rawResponse.Data.MarshalJSON()
		log.Println("signald-go: error unmarshalling response from signald of type", rawResponse.Type, string(rawResponseJson))
		return
	}

	return response, nil

}