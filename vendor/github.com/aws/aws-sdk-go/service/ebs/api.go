// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ebs

import (
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
)

const opGetSnapshotBlock = "GetSnapshotBlock"

// GetSnapshotBlockRequest generates a "aws/request.Request" representing the
// client's request for the GetSnapshotBlock operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetSnapshotBlock for more information on using the GetSnapshotBlock
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetSnapshotBlockRequest method.
//    req, resp := client.GetSnapshotBlockRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/ebs-2019-11-02/GetSnapshotBlock
func (c *EBS) GetSnapshotBlockRequest(input *GetSnapshotBlockInput) (req *request.Request, output *GetSnapshotBlockOutput) {
	op := &request.Operation{
		Name:       opGetSnapshotBlock,
		HTTPMethod: "GET",
		HTTPPath:   "/snapshots/{snapshotId}/blocks/{blockIndex}",
	}

	if input == nil {
		input = &GetSnapshotBlockInput{}
	}

	output = &GetSnapshotBlockOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetSnapshotBlock API operation for Amazon Elastic Block Store.
//
// Returns the data in a block in an Amazon Elastic Block Store snapshot.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Elastic Block Store's
// API operation GetSnapshotBlock for usage and error information.
//
// Returned Error Types:
//   * ValidationException
//   The input fails to satisfy the constraints of the EBS direct APIs.
//
//   * ResourceNotFoundException
//   The specified resource does not exist.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/ebs-2019-11-02/GetSnapshotBlock
func (c *EBS) GetSnapshotBlock(input *GetSnapshotBlockInput) (*GetSnapshotBlockOutput, error) {
	req, out := c.GetSnapshotBlockRequest(input)
	return out, req.Send()
}

// GetSnapshotBlockWithContext is the same as GetSnapshotBlock with the addition of
// the ability to pass a context and additional request options.
//
// See GetSnapshotBlock for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EBS) GetSnapshotBlockWithContext(ctx aws.Context, input *GetSnapshotBlockInput, opts ...request.Option) (*GetSnapshotBlockOutput, error) {
	req, out := c.GetSnapshotBlockRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListChangedBlocks = "ListChangedBlocks"

// ListChangedBlocksRequest generates a "aws/request.Request" representing the
// client's request for the ListChangedBlocks operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ListChangedBlocks for more information on using the ListChangedBlocks
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ListChangedBlocksRequest method.
//    req, resp := client.ListChangedBlocksRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/ebs-2019-11-02/ListChangedBlocks
func (c *EBS) ListChangedBlocksRequest(input *ListChangedBlocksInput) (req *request.Request, output *ListChangedBlocksOutput) {
	op := &request.Operation{
		Name:       opListChangedBlocks,
		HTTPMethod: "GET",
		HTTPPath:   "/snapshots/{secondSnapshotId}/changedblocks",
		Paginator: &request.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListChangedBlocksInput{}
	}

	output = &ListChangedBlocksOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListChangedBlocks API operation for Amazon Elastic Block Store.
//
// Returns the block indexes and block tokens for blocks that are different
// between two Amazon Elastic Block Store snapshots of the same volume/snapshot
// lineage.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Elastic Block Store's
// API operation ListChangedBlocks for usage and error information.
//
// Returned Error Types:
//   * ValidationException
//   The input fails to satisfy the constraints of the EBS direct APIs.
//
//   * ResourceNotFoundException
//   The specified resource does not exist.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/ebs-2019-11-02/ListChangedBlocks
func (c *EBS) ListChangedBlocks(input *ListChangedBlocksInput) (*ListChangedBlocksOutput, error) {
	req, out := c.ListChangedBlocksRequest(input)
	return out, req.Send()
}

// ListChangedBlocksWithContext is the same as ListChangedBlocks with the addition of
// the ability to pass a context and additional request options.
//
// See ListChangedBlocks for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EBS) ListChangedBlocksWithContext(ctx aws.Context, input *ListChangedBlocksInput, opts ...request.Option) (*ListChangedBlocksOutput, error) {
	req, out := c.ListChangedBlocksRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// ListChangedBlocksPages iterates over the pages of a ListChangedBlocks operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See ListChangedBlocks method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a ListChangedBlocks operation.
//    pageNum := 0
//    err := client.ListChangedBlocksPages(params,
//        func(page *ebs.ListChangedBlocksOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *EBS) ListChangedBlocksPages(input *ListChangedBlocksInput, fn func(*ListChangedBlocksOutput, bool) bool) error {
	return c.ListChangedBlocksPagesWithContext(aws.BackgroundContext(), input, fn)
}

// ListChangedBlocksPagesWithContext same as ListChangedBlocksPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EBS) ListChangedBlocksPagesWithContext(ctx aws.Context, input *ListChangedBlocksInput, fn func(*ListChangedBlocksOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *ListChangedBlocksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.ListChangedBlocksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	for p.Next() {
		if !fn(p.Page().(*ListChangedBlocksOutput), !p.HasNextPage()) {
			break
		}
	}

	return p.Err()
}

const opListSnapshotBlocks = "ListSnapshotBlocks"

// ListSnapshotBlocksRequest generates a "aws/request.Request" representing the
// client's request for the ListSnapshotBlocks operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ListSnapshotBlocks for more information on using the ListSnapshotBlocks
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ListSnapshotBlocksRequest method.
//    req, resp := client.ListSnapshotBlocksRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/ebs-2019-11-02/ListSnapshotBlocks
func (c *EBS) ListSnapshotBlocksRequest(input *ListSnapshotBlocksInput) (req *request.Request, output *ListSnapshotBlocksOutput) {
	op := &request.Operation{
		Name:       opListSnapshotBlocks,
		HTTPMethod: "GET",
		HTTPPath:   "/snapshots/{snapshotId}/blocks",
		Paginator: &request.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListSnapshotBlocksInput{}
	}

	output = &ListSnapshotBlocksOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListSnapshotBlocks API operation for Amazon Elastic Block Store.
//
// Returns the block indexes and block tokens for blocks in an Amazon Elastic
// Block Store snapshot.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Elastic Block Store's
// API operation ListSnapshotBlocks for usage and error information.
//
// Returned Error Types:
//   * ValidationException
//   The input fails to satisfy the constraints of the EBS direct APIs.
//
//   * ResourceNotFoundException
//   The specified resource does not exist.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/ebs-2019-11-02/ListSnapshotBlocks
func (c *EBS) ListSnapshotBlocks(input *ListSnapshotBlocksInput) (*ListSnapshotBlocksOutput, error) {
	req, out := c.ListSnapshotBlocksRequest(input)
	return out, req.Send()
}

// ListSnapshotBlocksWithContext is the same as ListSnapshotBlocks with the addition of
// the ability to pass a context and additional request options.
//
// See ListSnapshotBlocks for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EBS) ListSnapshotBlocksWithContext(ctx aws.Context, input *ListSnapshotBlocksInput, opts ...request.Option) (*ListSnapshotBlocksOutput, error) {
	req, out := c.ListSnapshotBlocksRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// ListSnapshotBlocksPages iterates over the pages of a ListSnapshotBlocks operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See ListSnapshotBlocks method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a ListSnapshotBlocks operation.
//    pageNum := 0
//    err := client.ListSnapshotBlocksPages(params,
//        func(page *ebs.ListSnapshotBlocksOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *EBS) ListSnapshotBlocksPages(input *ListSnapshotBlocksInput, fn func(*ListSnapshotBlocksOutput, bool) bool) error {
	return c.ListSnapshotBlocksPagesWithContext(aws.BackgroundContext(), input, fn)
}

// ListSnapshotBlocksPagesWithContext same as ListSnapshotBlocksPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EBS) ListSnapshotBlocksPagesWithContext(ctx aws.Context, input *ListSnapshotBlocksInput, fn func(*ListSnapshotBlocksOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *ListSnapshotBlocksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.ListSnapshotBlocksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	for p.Next() {
		if !fn(p.Page().(*ListSnapshotBlocksOutput), !p.HasNextPage()) {
			break
		}
	}

	return p.Err()
}

// A block of data in an Amazon Elastic Block Store snapshot.
type Block struct {
	_ struct{} `type:"structure"`

	// The block index.
	BlockIndex *int64 `type:"integer"`

	// The block token for the block index.
	BlockToken *string `type:"string"`
}

// String returns the string representation
func (s Block) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Block) GoString() string {
	return s.String()
}

// SetBlockIndex sets the BlockIndex field's value.
func (s *Block) SetBlockIndex(v int64) *Block {
	s.BlockIndex = &v
	return s
}

// SetBlockToken sets the BlockToken field's value.
func (s *Block) SetBlockToken(v string) *Block {
	s.BlockToken = &v
	return s
}

// A block of data in an Amazon Elastic Block Store snapshot that is different
// from another snapshot of the same volume/snapshot lineage.
type ChangedBlock struct {
	_ struct{} `type:"structure" sensitive:"true"`

	// The block index.
	BlockIndex *int64 `type:"integer"`

	// The block token for the block index of the FirstSnapshotId specified in the
	// ListChangedBlocks operation. This value is absent if the first snapshot does
	// not have the changed block that is on the second snapshot.
	FirstBlockToken *string `type:"string"`

	// The block token for the block index of the SecondSnapshotId specified in
	// the ListChangedBlocks operation.
	SecondBlockToken *string `type:"string"`
}

// String returns the string representation
func (s ChangedBlock) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ChangedBlock) GoString() string {
	return s.String()
}

// SetBlockIndex sets the BlockIndex field's value.
func (s *ChangedBlock) SetBlockIndex(v int64) *ChangedBlock {
	s.BlockIndex = &v
	return s
}

// SetFirstBlockToken sets the FirstBlockToken field's value.
func (s *ChangedBlock) SetFirstBlockToken(v string) *ChangedBlock {
	s.FirstBlockToken = &v
	return s
}

// SetSecondBlockToken sets the SecondBlockToken field's value.
func (s *ChangedBlock) SetSecondBlockToken(v string) *ChangedBlock {
	s.SecondBlockToken = &v
	return s
}

type GetSnapshotBlockInput struct {
	_ struct{} `type:"structure"`

	// The block index of the block from which to get data.
	//
	// Obtain the BlockIndex by running the ListChangedBlocks or ListSnapshotBlocks
	// operations.
	//
	// BlockIndex is a required field
	BlockIndex *int64 `location:"uri" locationName:"blockIndex" type:"integer" required:"true"`

	// The block token of the block from which to get data.
	//
	// Obtain the BlockToken by running the ListChangedBlocks or ListSnapshotBlocks
	// operations.
	//
	// BlockToken is a required field
	BlockToken *string `location:"querystring" locationName:"blockToken" type:"string" required:"true"`

	// The ID of the snapshot containing the block from which to get data.
	//
	// SnapshotId is a required field
	SnapshotId *string `location:"uri" locationName:"snapshotId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSnapshotBlockInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetSnapshotBlockInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSnapshotBlockInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetSnapshotBlockInput"}
	if s.BlockIndex == nil {
		invalidParams.Add(request.NewErrParamRequired("BlockIndex"))
	}
	if s.BlockToken == nil {
		invalidParams.Add(request.NewErrParamRequired("BlockToken"))
	}
	if s.SnapshotId == nil {
		invalidParams.Add(request.NewErrParamRequired("SnapshotId"))
	}
	if s.SnapshotId != nil && len(*s.SnapshotId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("SnapshotId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBlockIndex sets the BlockIndex field's value.
func (s *GetSnapshotBlockInput) SetBlockIndex(v int64) *GetSnapshotBlockInput {
	s.BlockIndex = &v
	return s
}

// SetBlockToken sets the BlockToken field's value.
func (s *GetSnapshotBlockInput) SetBlockToken(v string) *GetSnapshotBlockInput {
	s.BlockToken = &v
	return s
}

// SetSnapshotId sets the SnapshotId field's value.
func (s *GetSnapshotBlockInput) SetSnapshotId(v string) *GetSnapshotBlockInput {
	s.SnapshotId = &v
	return s
}

type GetSnapshotBlockOutput struct {
	_ struct{} `type:"structure" payload:"BlockData"`

	// The data content of the block.
	BlockData io.ReadCloser `type:"blob" sensitive:"true"`

	// The checksum generated for the block, which is Base64 encoded.
	Checksum *string `location:"header" locationName:"x-amz-Checksum" type:"string"`

	// The algorithm used to generate the checksum for the block, such as SHA256.
	ChecksumAlgorithm *string `location:"header" locationName:"x-amz-Checksum-Algorithm" type:"string" enum:"ChecksumAlgorithm"`

	// The size of the data in the block.
	DataLength *int64 `location:"header" locationName:"x-amz-Data-Length" type:"integer"`
}

// String returns the string representation
func (s GetSnapshotBlockOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetSnapshotBlockOutput) GoString() string {
	return s.String()
}

// SetBlockData sets the BlockData field's value.
func (s *GetSnapshotBlockOutput) SetBlockData(v io.ReadCloser) *GetSnapshotBlockOutput {
	s.BlockData = v
	return s
}

// SetChecksum sets the Checksum field's value.
func (s *GetSnapshotBlockOutput) SetChecksum(v string) *GetSnapshotBlockOutput {
	s.Checksum = &v
	return s
}

// SetChecksumAlgorithm sets the ChecksumAlgorithm field's value.
func (s *GetSnapshotBlockOutput) SetChecksumAlgorithm(v string) *GetSnapshotBlockOutput {
	s.ChecksumAlgorithm = &v
	return s
}

// SetDataLength sets the DataLength field's value.
func (s *GetSnapshotBlockOutput) SetDataLength(v int64) *GetSnapshotBlockOutput {
	s.DataLength = &v
	return s
}

type ListChangedBlocksInput struct {
	_ struct{} `type:"structure"`

	// The ID of the first snapshot to use for the comparison.
	//
	// The FirstSnapshotID parameter must be specified with a SecondSnapshotId parameter;
	// otherwise, an error occurs.
	FirstSnapshotId *string `location:"querystring" locationName:"firstSnapshotId" min:"1" type:"string"`

	// The number of results to return.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"100" type:"integer"`

	// The token to request the next page of results.
	NextToken *string `location:"querystring" locationName:"pageToken" type:"string"`

	// The ID of the second snapshot to use for the comparison.
	//
	// The SecondSnapshotId parameter must be specified with a FirstSnapshotID parameter;
	// otherwise, an error occurs.
	//
	// SecondSnapshotId is a required field
	SecondSnapshotId *string `location:"uri" locationName:"secondSnapshotId" min:"1" type:"string" required:"true"`

	// The block index from which the comparison should start.
	//
	// The list in the response will start from this block index or the next valid
	// block index in the snapshots.
	StartingBlockIndex *int64 `location:"querystring" locationName:"startingBlockIndex" type:"integer"`
}

// String returns the string representation
func (s ListChangedBlocksInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListChangedBlocksInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListChangedBlocksInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListChangedBlocksInput"}
	if s.FirstSnapshotId != nil && len(*s.FirstSnapshotId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("FirstSnapshotId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 100 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 100))
	}
	if s.SecondSnapshotId == nil {
		invalidParams.Add(request.NewErrParamRequired("SecondSnapshotId"))
	}
	if s.SecondSnapshotId != nil && len(*s.SecondSnapshotId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("SecondSnapshotId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFirstSnapshotId sets the FirstSnapshotId field's value.
func (s *ListChangedBlocksInput) SetFirstSnapshotId(v string) *ListChangedBlocksInput {
	s.FirstSnapshotId = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListChangedBlocksInput) SetMaxResults(v int64) *ListChangedBlocksInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListChangedBlocksInput) SetNextToken(v string) *ListChangedBlocksInput {
	s.NextToken = &v
	return s
}

// SetSecondSnapshotId sets the SecondSnapshotId field's value.
func (s *ListChangedBlocksInput) SetSecondSnapshotId(v string) *ListChangedBlocksInput {
	s.SecondSnapshotId = &v
	return s
}

// SetStartingBlockIndex sets the StartingBlockIndex field's value.
func (s *ListChangedBlocksInput) SetStartingBlockIndex(v int64) *ListChangedBlocksInput {
	s.StartingBlockIndex = &v
	return s
}

type ListChangedBlocksOutput struct {
	_ struct{} `type:"structure"`

	// The size of the block.
	BlockSize *int64 `type:"integer"`

	// An array of objects containing information about the changed blocks.
	ChangedBlocks []*ChangedBlock `type:"list"`

	// The time when the BlockToken expires.
	ExpiryTime *time.Time `type:"timestamp"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `type:"string"`

	// The size of the volume in GB.
	VolumeSize *int64 `type:"long"`
}

// String returns the string representation
func (s ListChangedBlocksOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListChangedBlocksOutput) GoString() string {
	return s.String()
}

// SetBlockSize sets the BlockSize field's value.
func (s *ListChangedBlocksOutput) SetBlockSize(v int64) *ListChangedBlocksOutput {
	s.BlockSize = &v
	return s
}

// SetChangedBlocks sets the ChangedBlocks field's value.
func (s *ListChangedBlocksOutput) SetChangedBlocks(v []*ChangedBlock) *ListChangedBlocksOutput {
	s.ChangedBlocks = v
	return s
}

// SetExpiryTime sets the ExpiryTime field's value.
func (s *ListChangedBlocksOutput) SetExpiryTime(v time.Time) *ListChangedBlocksOutput {
	s.ExpiryTime = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListChangedBlocksOutput) SetNextToken(v string) *ListChangedBlocksOutput {
	s.NextToken = &v
	return s
}

// SetVolumeSize sets the VolumeSize field's value.
func (s *ListChangedBlocksOutput) SetVolumeSize(v int64) *ListChangedBlocksOutput {
	s.VolumeSize = &v
	return s
}

type ListSnapshotBlocksInput struct {
	_ struct{} `type:"structure"`

	// The number of results to return.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"100" type:"integer"`

	// The token to request the next page of results.
	NextToken *string `location:"querystring" locationName:"pageToken" type:"string"`

	// The ID of the snapshot from which to get block indexes and block tokens.
	//
	// SnapshotId is a required field
	SnapshotId *string `location:"uri" locationName:"snapshotId" min:"1" type:"string" required:"true"`

	// The block index from which the list should start. The list in the response
	// will start from this block index or the next valid block index in the snapshot.
	StartingBlockIndex *int64 `location:"querystring" locationName:"startingBlockIndex" type:"integer"`
}

// String returns the string representation
func (s ListSnapshotBlocksInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListSnapshotBlocksInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSnapshotBlocksInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListSnapshotBlocksInput"}
	if s.MaxResults != nil && *s.MaxResults < 100 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 100))
	}
	if s.SnapshotId == nil {
		invalidParams.Add(request.NewErrParamRequired("SnapshotId"))
	}
	if s.SnapshotId != nil && len(*s.SnapshotId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("SnapshotId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListSnapshotBlocksInput) SetMaxResults(v int64) *ListSnapshotBlocksInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListSnapshotBlocksInput) SetNextToken(v string) *ListSnapshotBlocksInput {
	s.NextToken = &v
	return s
}

// SetSnapshotId sets the SnapshotId field's value.
func (s *ListSnapshotBlocksInput) SetSnapshotId(v string) *ListSnapshotBlocksInput {
	s.SnapshotId = &v
	return s
}

// SetStartingBlockIndex sets the StartingBlockIndex field's value.
func (s *ListSnapshotBlocksInput) SetStartingBlockIndex(v int64) *ListSnapshotBlocksInput {
	s.StartingBlockIndex = &v
	return s
}

type ListSnapshotBlocksOutput struct {
	_ struct{} `type:"structure"`

	// The size of the block.
	BlockSize *int64 `type:"integer"`

	// An array of objects containing information about the blocks.
	Blocks []*Block `type:"list" sensitive:"true"`

	// The time when the BlockToken expires.
	ExpiryTime *time.Time `type:"timestamp"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `type:"string"`

	// The size of the volume in GB.
	VolumeSize *int64 `type:"long"`
}

// String returns the string representation
func (s ListSnapshotBlocksOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListSnapshotBlocksOutput) GoString() string {
	return s.String()
}

// SetBlockSize sets the BlockSize field's value.
func (s *ListSnapshotBlocksOutput) SetBlockSize(v int64) *ListSnapshotBlocksOutput {
	s.BlockSize = &v
	return s
}

// SetBlocks sets the Blocks field's value.
func (s *ListSnapshotBlocksOutput) SetBlocks(v []*Block) *ListSnapshotBlocksOutput {
	s.Blocks = v
	return s
}

// SetExpiryTime sets the ExpiryTime field's value.
func (s *ListSnapshotBlocksOutput) SetExpiryTime(v time.Time) *ListSnapshotBlocksOutput {
	s.ExpiryTime = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListSnapshotBlocksOutput) SetNextToken(v string) *ListSnapshotBlocksOutput {
	s.NextToken = &v
	return s
}

// SetVolumeSize sets the VolumeSize field's value.
func (s *ListSnapshotBlocksOutput) SetVolumeSize(v int64) *ListSnapshotBlocksOutput {
	s.VolumeSize = &v
	return s
}

// The specified resource does not exist.
type ResourceNotFoundException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation
func (s ResourceNotFoundException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceNotFoundException) GoString() string {
	return s.String()
}

func newErrorResourceNotFoundException(v protocol.ResponseMetadata) error {
	return &ResourceNotFoundException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ResourceNotFoundException) Code() string {
	return "ResourceNotFoundException"
}

// Message returns the exception's message.
func (s *ResourceNotFoundException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ResourceNotFoundException) OrigErr() error {
	return nil
}

func (s *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ResourceNotFoundException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ResourceNotFoundException) RequestID() string {
	return s.RespMetadata.RequestID
}

// The input fails to satisfy the constraints of the EBS direct APIs.
type ValidationException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`

	// The reason for the validation exception.
	Reason *string `type:"string" enum:"ValidationExceptionReason"`
}

// String returns the string representation
func (s ValidationException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ValidationException) GoString() string {
	return s.String()
}

func newErrorValidationException(v protocol.ResponseMetadata) error {
	return &ValidationException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ValidationException) Code() string {
	return "ValidationException"
}

// Message returns the exception's message.
func (s *ValidationException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ValidationException) OrigErr() error {
	return nil
}

func (s *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s\n%s", s.Code(), s.Message(), s.String())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ValidationException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ValidationException) RequestID() string {
	return s.RespMetadata.RequestID
}

const (
	// ChecksumAlgorithmSha256 is a ChecksumAlgorithm enum value
	ChecksumAlgorithmSha256 = "SHA256"
)

const (
	// ValidationExceptionReasonInvalidCustomerKey is a ValidationExceptionReason enum value
	ValidationExceptionReasonInvalidCustomerKey = "INVALID_CUSTOMER_KEY"

	// ValidationExceptionReasonInvalidPageToken is a ValidationExceptionReason enum value
	ValidationExceptionReasonInvalidPageToken = "INVALID_PAGE_TOKEN"

	// ValidationExceptionReasonInvalidBlockToken is a ValidationExceptionReason enum value
	ValidationExceptionReasonInvalidBlockToken = "INVALID_BLOCK_TOKEN"

	// ValidationExceptionReasonInvalidSnapshotId is a ValidationExceptionReason enum value
	ValidationExceptionReasonInvalidSnapshotId = "INVALID_SNAPSHOT_ID"

	// ValidationExceptionReasonUnrelatedSnapshots is a ValidationExceptionReason enum value
	ValidationExceptionReasonUnrelatedSnapshots = "UNRELATED_SNAPSHOTS"
)
