package v3action_test

import (
	"errors"
	"net/url"

	. "code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/actor/v3action/v3actionfakes"
	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Isolation Segment Actions", func() {
	var (
		actor                     Actor
		fakeCloudControllerClient *v3actionfakes.FakeCloudControllerClient
	)

	BeforeEach(func() {
		fakeCloudControllerClient = new(v3actionfakes.FakeCloudControllerClient)
		actor = NewActor(fakeCloudControllerClient)
	})

	Describe("CreateIsolationSegment", func() {
		Context("when the create is successful", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.CreateIsolationSegmentReturns(
					ccv3.IsolationSegment{},
					ccv3.Warnings{"warning-1", "warning-2"},
					nil,
				)
			})

			It("returns all warnings", func() {
				warnings, err := actor.CreateIsolationSegmentByName("some-isolation-segment-guid")
				Expect(err).ToNot(HaveOccurred())

				Expect(warnings).To(ConsistOf("warning-1", "warning-2"))

				Expect(fakeCloudControllerClient.CreateIsolationSegmentCallCount()).To(Equal(1))
				isolationSegmentName := fakeCloudControllerClient.CreateIsolationSegmentArgsForCall(0)
				Expect(isolationSegmentName).To(Equal("some-isolation-segment-guid"))
			})
		})

		Context("when the cloud controller client returns an error", func() {
			Context("when an unexpected error occurs", func() {
				var expectedErr error

				BeforeEach(func() {
					expectedErr = errors.New("I am a CloudControllerClient Error")
					fakeCloudControllerClient.CreateIsolationSegmentReturns(
						ccv3.IsolationSegment{},
						ccv3.Warnings{"warning-1", "warning-2"},
						expectedErr,
					)
				})

				It("returns the same error and all warnings", func() {
					warnings, err := actor.CreateIsolationSegmentByName("isolation-segment")
					Expect(err).To(MatchError(expectedErr))
					Expect(warnings).To(ConsistOf("warning-1", "warning-2"))
				})
			})

			Context("when an UnprocessableEntityError occurs", func() {
				BeforeEach(func() {
					fakeCloudControllerClient.CreateIsolationSegmentReturns(
						ccv3.IsolationSegment{},
						ccv3.Warnings{"warning-1", "warning-2"},
						cloudcontroller.UnprocessableEntityError{},
					)
				})

				It("returns an IsolationSegmentAlreadyExistsError and all warnings", func() {
					warnings, err := actor.CreateIsolationSegmentByName("isolation-segment")
					Expect(err).To(MatchError(IsolationSegmentAlreadyExistsError{Name: "isolation-segment"}))
					Expect(warnings).To(ConsistOf("warning-1", "warning-2"))
				})
			})
		})
	})

	Describe("DeleteIsolationSegmentByName", func() {
		Context("when the isolation segment is found", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.GetIsolationSegmentsReturns([]ccv3.IsolationSegment{
					{
						GUID: "some-iso-guid",
						Name: "some-iso-seg",
					},
				}, ccv3.Warnings{"I r warnings", "I are two warnings"},
					nil,
				)
			})

			Context("when the delete is successful", func() {
				BeforeEach(func() {
					fakeCloudControllerClient.DeleteIsolationSegmentReturns(ccv3.Warnings{"delete warning-1", "delete warning-2"}, nil)
				})

				It("returns back all warnings", func() {
					warnings, err := actor.DeleteIsolationSegmentByName("some-iso-seg")
					Expect(err).NotTo(HaveOccurred())
					Expect(warnings).To(ConsistOf("I r warnings", "I are two warnings", "delete warning-1", "delete warning-2"))

					Expect(fakeCloudControllerClient.GetIsolationSegmentsCallCount()).To(Equal(1))
					Expect(fakeCloudControllerClient.GetIsolationSegmentsArgsForCall(0)).To(Equal(url.Values{ccv3.NameFilter: []string{"some-iso-seg"}}))

					Expect(fakeCloudControllerClient.DeleteIsolationSegmentCallCount()).To(Equal(1))
					Expect(fakeCloudControllerClient.DeleteIsolationSegmentArgsForCall(0)).To(Equal("some-iso-guid"))
				})
			})

			Context("when the delete returns an error", func() {
				var expectedErr error

				BeforeEach(func() {
					expectedErr = errors.New("some-cc-error")
					fakeCloudControllerClient.DeleteIsolationSegmentReturns(ccv3.Warnings{"delete warning-1", "delete warning-2"}, expectedErr)
				})

				It("returns back the error and all warnings", func() {
					warnings, err := actor.DeleteIsolationSegmentByName("some-iso-seg")
					Expect(warnings).To(ConsistOf("I r warnings", "I are two warnings", "delete warning-1", "delete warning-2"))
					Expect(err).To(MatchError(expectedErr))
				})
			})
		})

		Context("when the search errors", func() {
			var expectedErr error

			BeforeEach(func() {
				expectedErr = errors.New("some-cc-error")
				fakeCloudControllerClient.GetIsolationSegmentsReturns(nil, ccv3.Warnings{"I r warnings", "I are two warnings"}, expectedErr)
			})

			It("returns the error and all warnings", func() {
				warnings, err := actor.DeleteIsolationSegmentByName("some-iso-seg")
				Expect(warnings).To(ConsistOf("I r warnings", "I are two warnings"))
				Expect(err).To(MatchError(expectedErr))
			})
		})
	})

	Describe("GetIsolationSegmentByName", func() {
		Context("when the isolation segment exists", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.GetIsolationSegmentsReturns([]ccv3.IsolationSegment{
					{
						GUID: "some-iso-guid",
						Name: "some-iso-seg",
					},
				}, ccv3.Warnings{"I r warnings", "I are two warnings"},
					nil,
				)
			})

			It("returns the isolation segment and warnings", func() {
				segment, warnings, err := actor.GetIsolationSegmentByName("some-iso-seg")
				Expect(err).NotTo(HaveOccurred())
				Expect(warnings).To(ConsistOf("I r warnings", "I are two warnings"))
				Expect(segment).To(Equal(IsolationSegment{
					GUID: "some-iso-guid",
					Name: "some-iso-seg",
				}))

				Expect(fakeCloudControllerClient.GetIsolationSegmentsCallCount()).To(Equal(1))
				Expect(fakeCloudControllerClient.GetIsolationSegmentsArgsForCall(0)).To(Equal(url.Values{ccv3.NameFilter: []string{"some-iso-seg"}}))
			})
		})

		Context("when the isolation segment does *not* exist", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.GetIsolationSegmentsReturns(nil, ccv3.Warnings{"I r warnings", "I are two warnings"}, nil)
			})

			It("returns an IsolationSegmentNotFoundError", func() {
				_, warnings, err := actor.GetIsolationSegmentByName("some-iso-seg")
				Expect(err).To(MatchError(IsolationSegmentNotFoundError{Name: "some-iso-seg"}))
				Expect(warnings).To(ConsistOf("I r warnings", "I are two warnings"))
			})
		})

		Context("when the cloud controller errors", func() {
			var expectedErr error
			BeforeEach(func() {
				expectedErr = errors.New("some-cc-error")
				fakeCloudControllerClient.GetIsolationSegmentsReturns(nil, ccv3.Warnings{"I r warnings", "I are two warnings"}, expectedErr)
			})

			It("returns the error and all warnings", func() {
				_, warnings, err := actor.GetIsolationSegmentByName("some-iso-seg")
				Expect(err).To(MatchError(expectedErr))
				Expect(warnings).To(ConsistOf("I r warnings", "I are two warnings"))
			})
		})
	})
})
