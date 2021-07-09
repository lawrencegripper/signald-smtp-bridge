// Copyright © 2021 Finn Herzfeld <finn@janky.solutions>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package v0

// Request represents a message sent to signald
type LegacyRequest struct {
	Type                string           `json:"type"`
	ID                  string           `json:"id,omitempty"`
	Username            string           `json:"username,omitempty"`
	MessageBody         string           `json:"messageBody,omitempty"`
	RecipientAddress    *JsonAddress     `json:"recipientAddress,omitempty"`
	RecipientGroupID    string           `json:"recipientGroupId,omitempty"`
	Voice               bool             `json:"voice,omitempty"`
	Code                string           `json:"code,omitempty"`
	DeviceName          string           `json:"deviceName,omitempty"`
	AttachmentFilenames []string         `json:"attachmentFilenames,omitempty"`
	URI                 string           `json:"uri,omitempty"`
	Attachments         []JsonAttachment `json:"attachments,omitempty"`
	GroupName           string           `json:"groupName,omitempty"`
	Members             []string         `json:"members,omitempty"`
	Avatar              string           `json:"avatar,omitempty"`
	Captcha             string           `json:"captcha,omitempty"`
	Version             string           `json:"version,omitempty"`
}
