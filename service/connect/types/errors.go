// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// You do not have sufficient permissions to perform this action.
type AccessDeniedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The contact flow has not been published.
type ContactFlowNotPublishedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ContactFlowNotPublishedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ContactFlowNotPublishedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ContactFlowNotPublishedException) ErrorCode() string {
	return "ContactFlowNotPublishedException"
}
func (e *ContactFlowNotPublishedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The contact with the specified ID is not active or does not exist.
type ContactNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ContactNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ContactNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ContactNotFoundException) ErrorCode() string             { return "ContactNotFoundException" }
func (e *ContactNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Outbound calls to the destination number are not allowed.
type DestinationNotAllowedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DestinationNotAllowedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DestinationNotAllowedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DestinationNotAllowedException) ErrorCode() string             { return "DestinationNotAllowedException" }
func (e *DestinationNotAllowedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource with the specified name already exists.
type DuplicateResourceException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DuplicateResourceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateResourceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateResourceException) ErrorCode() string             { return "DuplicateResourceException" }
func (e *DuplicateResourceException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An entity with the same name already exists.
type IdempotencyException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *IdempotencyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IdempotencyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IdempotencyException) ErrorCode() string             { return "IdempotencyException" }
func (e *IdempotencyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Request processing failed because of an error or failure with the service.
type InternalServiceException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InternalServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceException) ErrorCode() string             { return "InternalServiceException" }
func (e *InternalServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The contact flow is not valid.
type InvalidContactFlowException struct {
	Message *string

	Problems []ProblemDetail

	noSmithyDocumentSerde
}

func (e *InvalidContactFlowException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidContactFlowException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidContactFlowException) ErrorCode() string             { return "InvalidContactFlowException" }
func (e *InvalidContactFlowException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The problems with the module. Please fix before trying again.
type InvalidContactFlowModuleException struct {
	Message *string

	Problems []ProblemDetail

	noSmithyDocumentSerde
}

func (e *InvalidContactFlowModuleException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidContactFlowModuleException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidContactFlowModuleException) ErrorCode() string {
	return "InvalidContactFlowModuleException"
}
func (e *InvalidContactFlowModuleException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more of the specified parameters are not valid.
type InvalidParameterException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string             { return "InvalidParameterException" }
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request is not valid.
type InvalidRequestException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string             { return "InvalidRequestException" }
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The allowed limit for the resource has been exceeded.
type LimitExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The contact is not permitted.
type OutboundContactNotPermittedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *OutboundContactNotPermittedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OutboundContactNotPermittedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OutboundContactNotPermittedException) ErrorCode() string {
	return "OutboundContactNotPermittedException"
}
func (e *OutboundContactNotPermittedException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

type PropertyValidationException struct {
	Message *string

	PropertyList []PropertyValidationExceptionProperty

	noSmithyDocumentSerde
}

func (e *PropertyValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PropertyValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PropertyValidationException) ErrorCode() string             { return "PropertyValidationException" }
func (e *PropertyValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource already has that name.
type ResourceConflictException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceConflictException) ErrorCode() string             { return "ResourceConflictException" }
func (e *ResourceConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// That resource is already in use. Please try another.
type ResourceInUseException struct {
	Message *string

	ResourceType ResourceType
	ResourceId   *string

	noSmithyDocumentSerde
}

func (e *ResourceInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUseException) ErrorCode() string             { return "ResourceInUseException" }
func (e *ResourceInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource was not found.
type ResourceNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service quota has been exceeded.
type ServiceQuotaExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ServiceQuotaExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceQuotaExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceQuotaExceededException) ErrorCode() string             { return "ServiceQuotaExceededException" }
func (e *ServiceQuotaExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The throttling limit has been exceeded.
type ThrottlingException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string             { return "ThrottlingException" }
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// No user with the specified credentials was found in the Amazon Connect instance.
type UserNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UserNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UserNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UserNotFoundException) ErrorCode() string             { return "UserNotFoundException" }
func (e *UserNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
