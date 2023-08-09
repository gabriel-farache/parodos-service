package workflows_test

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/parodos-dev/parodos-service/pkg/errors"
	"github.com/parodos-dev/parodos-service/pkg/workflows"
)

type CustomResponseWriter struct {
	body       []byte
	statusCode int
	header     http.Header
}

func NewCustomResponseWriter() *CustomResponseWriter {
	return &CustomResponseWriter{
		header: http.Header{},
	}
}

func (w *CustomResponseWriter) Header() http.Header {
	return w.header
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.body = b
	// implement it as per your requirement
	return 0, nil
}

func (w *CustomResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

var _ = Describe("Backend factory", func() {

	var (
		mockCtrl *gomock.Controller
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
	})

	Context("Group handler", func() {
		It("should get all groups", func() {
			// given
			w := NewCustomResponseWriter()
			ctx, _ := gin.CreateTestContext(w)
			groups := []workflows.Group{
				{Name: "testGroup01", Repository: "repoTest01"},
				{Name: "testGroup02", Repository: "repoTest02"},
			}
			expectedGroups, err := json.Marshal(groups)
			if err != nil {
				fmt.Println(err)
				return
			}
			mockQuery := workflows.NewMockWorkflowsQuery(mockCtrl)
			mockQuery.EXPECT().ListGroups().Times(1).Return(groups, nil)
			handler := workflows.NewWorkflowDefinitionHandler(&workflows.WorkflowCommandHandler{}, mockQuery)
			// when
			handler.GetGroups(ctx)
			// then
			Expect(w.statusCode).To(Equal(http.StatusOK))
			Expect(w.body).To(Equal(expectedGroups))
		})

		It("should error when get all groups", func() {
			// given
			w := NewCustomResponseWriter()
			ctx, _ := gin.CreateTestContext(w)
			groups := []workflows.Group{
				{Name: "testGroup01", Repository: "repoTest01"},
				{Name: "testGroup02", Repository: "repoTest02"},
			}
			expectedGroups, err := json.Marshal(groups)
			if err != nil {
				fmt.Println(err)
				return
			}
			mockQuery := workflows.NewMockWorkflowsQuery(mockCtrl)
			mockQuery.EXPECT().ListGroups().Times(1).Return(groups, fmt.Errorf("Error while getting all groups"))
			handler := workflows.NewWorkflowDefinitionHandler(&workflows.WorkflowCommandHandler{}, mockQuery)
			// when
			handler.GetGroups(ctx)
			// then
			Expect(w.statusCode).To(Equal(http.StatusInternalServerError))
			Expect(w.body).ToNot(Equal(expectedGroups))
		})

		It("should get given group", func() {
			// given
			w := NewCustomResponseWriter()
			ctx, _ := gin.CreateTestContext(w)
			groupID := "testGroup01"
			ctx.Set(workflows.GroupIdParam, groupID)
			group := &workflows.GroupDetails{
				workflows.Group{Name: "testGroup01", Repository: "repoTest01"},
				[]workflows.Workflow{
					{Name: "testWorkflow"},
				},
			}
			expectedGroup, err := json.Marshal(group)
			if err != nil {
				fmt.Println(err)
				return
			}
			mockQuery := workflows.NewMockWorkflowsQuery(mockCtrl)
			mockQuery.EXPECT().GetGroup(groupID).Times(1).Return(group, nil)
			handler := workflows.NewWorkflowDefinitionHandler(&workflows.WorkflowCommandHandler{}, mockQuery)
			// when
			handler.GetGroup(ctx)
			// then
			Expect(w.statusCode).To(Equal(http.StatusOK))
			Expect(w.body).To(Equal(expectedGroup))
		})

		It("should error as no group id in context when retreiving group", func() {
			// given
			w := NewCustomResponseWriter()
			ctx, _ := gin.CreateTestContext(w)
			group := workflows.GroupDetails{
				workflows.Group{Name: "testGroup01", Repository: "repoTest01"},
				[]workflows.Workflow{
					{Name: "testWorkflow"},
				},
			}
			expectedGroup, err := json.Marshal(group)
			if err != nil {
				fmt.Println(err)
				return
			}
			mockQuery := workflows.NewMockWorkflowsQuery(mockCtrl)
			mockQuery.EXPECT().GetGroup(gomock.Any()).Times(0)
			handler := workflows.NewWorkflowDefinitionHandler(&workflows.WorkflowCommandHandler{}, mockQuery)
			// when
			handler.GetGroup(ctx)
			// then
			Expect(w.statusCode).To(Equal(http.StatusBadRequest))
			Expect(w.body).ToNot(Equal(expectedGroup))
		})

		It("should error as group id does not exist when retreiving group", func() {
			// given
			w := NewCustomResponseWriter()
			ctx, _ := gin.CreateTestContext(w)
			group := workflows.GroupDetails{
				workflows.Group{Name: "testGroup01", Repository: "repoTest01"},
				[]workflows.Workflow{
					{Name: "testWorkflow"},
				},
			}
			expectedGroup, err := json.Marshal(group)
			if err != nil {
				fmt.Println(err)
				return
			}
			mockQuery := workflows.NewMockWorkflowsQuery(mockCtrl)
			mockQuery.EXPECT().GetGroup(gomock.Any()).Times(0)
			handler := workflows.NewWorkflowDefinitionHandler(&workflows.WorkflowCommandHandler{}, mockQuery)
			// when
			handler.GetGroup(ctx)
			// then
			Expect(w.statusCode).To(Equal(http.StatusBadRequest))
			Expect(w.body).ToNot(Equal(expectedGroup))
		})

		It("should error as empty group id given when retreiving group", func() {
			// given
			w := NewCustomResponseWriter()
			ctx, _ := gin.CreateTestContext(w)
			groupID := ""
			ctx.Set(workflows.GroupIdParam, groupID)
			group := workflows.GroupDetails{
				workflows.Group{Name: "testGroup01", Repository: "repoTest01"},
				[]workflows.Workflow{
					{Name: "testWorkflow"},
				},
			}
			expectedGroup, err := json.Marshal(group)
			if err != nil {
				fmt.Println(err)
				return
			}
			mockQuery := workflows.NewMockWorkflowsQuery(mockCtrl)
			mockQuery.EXPECT().GetGroup(groupID).Times(1).Return(nil, errors.BadRequestError{Message: "No group provided"})
			handler := workflows.NewWorkflowDefinitionHandler(&workflows.WorkflowCommandHandler{}, mockQuery)
			// when
			handler.GetGroup(ctx)
			// then
			Expect(w.statusCode).To(Equal(http.StatusBadRequest))
			Expect(w.body).ToNot(Equal(expectedGroup))
		})
	})
	Context("Workflow handler", func() {
		It("should get all workflows", func() {
			// given
			groupID := "testGroup01"
			w := NewCustomResponseWriter()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Set(workflows.GroupIdParam, groupID)
			workflowsList := []workflows.Workflow{
				{Name: "testW01"},
				{Name: "testW02"},
			}
			expectedWorkflows, err := json.Marshal(workflowsList)
			if err != nil {
				fmt.Println(err)
				return
			}
			mockQuery := workflows.NewMockWorkflowsQuery(mockCtrl)
			mockQuery.EXPECT().ListWorkflows(groupID).Times(1).Return(workflowsList, nil)
			handler := workflows.NewWorkflowDefinitionHandler(&workflows.WorkflowCommandHandler{}, mockQuery)
			// when
			handler.GetWorkflows(ctx)
			// then
			Expect(w.statusCode).To(Equal(http.StatusOK))
			Expect(w.body).To(Equal(expectedWorkflows))
		})

		It("should error when no group specified when get all workflows", func() {
			// given
			w := NewCustomResponseWriter()
			ctx, _ := gin.CreateTestContext(w)
			workflowsList := []workflows.Workflow{
				{Name: "testW01"},
				{Name: "testW02"},
			}
			expectedWorkflows, err := json.Marshal(workflowsList)
			if err != nil {
				fmt.Println(err)
				return
			}
			mockQuery := workflows.NewMockWorkflowsQuery(mockCtrl)
			mockQuery.EXPECT().ListWorkflows(gomock.Any()).Times(0)
			handler := workflows.NewWorkflowDefinitionHandler(&workflows.WorkflowCommandHandler{}, mockQuery)
			// when
			handler.GetWorkflows(ctx)
			// then
			Expect(w.statusCode).To(Equal(http.StatusBadRequest))
			Expect(w.body).ToNot(Equal(expectedWorkflows))
		})

		It("should error when get all workflows", func() {
			// given
			w := NewCustomResponseWriter()
			ctx, _ := gin.CreateTestContext(w)
			groupId := "testGroup01"
			ctx.Set(workflows.GroupIdParam, groupId)
			workflowsList := []workflows.Workflow{
				{Name: "testW01"},
				{Name: "testW02"},
			}
			expectedWorkflows, err := json.Marshal(workflowsList)
			if err != nil {
				fmt.Println(err)
				return
			}
			mockQuery := workflows.NewMockWorkflowsQuery(mockCtrl)
			mockQuery.EXPECT().ListWorkflows(groupId).Times(1).Return(workflowsList, fmt.Errorf("error get all workflows"))
			handler := workflows.NewWorkflowDefinitionHandler(&workflows.WorkflowCommandHandler{}, mockQuery)
			// when
			handler.GetWorkflows(ctx)
			// then
			Expect(w.statusCode).To(Equal(http.StatusInternalServerError))
			Expect(w.body).ToNot(Equal(expectedWorkflows))
		})

		It("should get given workflow", func() {
			// given
			w := NewCustomResponseWriter()
			ctx, _ := gin.CreateTestContext(w)
			groupId := "testGroup01"
			workflowID := "testW01"
			ctx.Set(workflows.GroupIdParam, groupId)
			ctx.Set(workflows.WorkflowIdParam, workflowID)
			workflow := &workflows.Workflow{Name: "testW01"}
			expectedWorkflow, err := json.Marshal(workflow)
			if err != nil {
				fmt.Println(err)
				return
			}
			mockQuery := workflows.NewMockWorkflowsQuery(mockCtrl)
			mockQuery.EXPECT().GetWorkflow(groupId, workflowID).Times(1).Return(workflow, nil)
			handler := workflows.NewWorkflowDefinitionHandler(&workflows.WorkflowCommandHandler{}, mockQuery)
			// when
			handler.GetWorkflow(ctx)
			// then
			Expect(w.statusCode).To(Equal(http.StatusOK))
			Expect(w.body).To(Equal(expectedWorkflow))
		})

		It("should error as no group id in context when retreiving workflow", func() {
			// given
			w := NewCustomResponseWriter()
			ctx, _ := gin.CreateTestContext(w)
			workflow := workflows.Workflow{Name: "testW01"}
			expectedWorkflow, err := json.Marshal(workflow)
			if err != nil {
				fmt.Println(err)
				return
			}

			mockQuery := workflows.NewMockWorkflowsQuery(mockCtrl)
			mockQuery.EXPECT().GetWorkflow(gomock.Any(), gomock.Any()).Times(0)
			handler := workflows.NewWorkflowDefinitionHandler(&workflows.WorkflowCommandHandler{}, mockQuery)
			// when
			handler.GetWorkflow(ctx)
			// then
			Expect(w.statusCode).To(Equal(http.StatusBadRequest))
			Expect(w.body).ToNot(Equal(expectedWorkflow))
		})

		It("should error as no workflow id in context when retreiving workflow", func() {
			// given
			w := NewCustomResponseWriter()
			ctx, _ := gin.CreateTestContext(w)
			groupId := "testGroup01"
			ctx.Set(workflows.GroupIdParam, groupId)
			workflow := workflows.Workflow{Name: "testW01"}
			expectedWorkflow, err := json.Marshal(workflow)
			if err != nil {
				fmt.Println(err)
				return
			}

			mockQuery := workflows.NewMockWorkflowsQuery(mockCtrl)
			mockQuery.EXPECT().GetWorkflow(groupId, gomock.Any()).Times(0)
			handler := workflows.NewWorkflowDefinitionHandler(&workflows.WorkflowCommandHandler{}, mockQuery)
			// when
			handler.GetWorkflow(ctx)
			// then
			Expect(w.statusCode).To(Equal(http.StatusBadRequest))
			Expect(w.body).ToNot(Equal(expectedWorkflow))
		})

		It("should error as empty group id given when retreiving workflow", func() {
			// given
			w := NewCustomResponseWriter()
			ctx, _ := gin.CreateTestContext(w)
			groupId := ""
			workflowID := ""
			ctx.Set(workflows.GroupIdParam, groupId)
			ctx.Set(workflows.WorkflowIdParam, workflowID)
			workflow := workflows.Workflow{Name: "testW01"}
			expectedWorkflow, err := json.Marshal(workflow)
			if err != nil {
				fmt.Println(err)
				return
			}

			mockQuery := workflows.NewMockWorkflowsQuery(mockCtrl)
			mockQuery.EXPECT().GetWorkflow(groupId, gomock.Any()).Times(1).Return(nil, errors.BadRequestError{Message: "not group provided"})
			handler := workflows.NewWorkflowDefinitionHandler(&workflows.WorkflowCommandHandler{}, mockQuery)
			// when
			handler.GetWorkflow(ctx)
			// then
			Expect(w.statusCode).To(Equal(http.StatusBadRequest))
			Expect(w.body).ToNot(Equal(expectedWorkflow))
		})

		It("should error as empty workflow id given when retreiving workflow", func() {
			// given
			w := NewCustomResponseWriter()
			ctx, _ := gin.CreateTestContext(w)
			groupId := "testGroup01"
			workflowID := ""
			ctx.Set(workflows.GroupIdParam, groupId)
			ctx.Set(workflows.WorkflowIdParam, workflowID)
			workflow := workflows.Workflow{Name: "testW01"}
			expectedWorkflow, err := json.Marshal(workflow)
			if err != nil {
				fmt.Println(err)
				return
			}

			mockQuery := workflows.NewMockWorkflowsQuery(mockCtrl)
			mockQuery.EXPECT().GetWorkflow(groupId, workflowID).Times(1).Return(nil, errors.BadRequestError{Message: "not workflow provided"})
			handler := workflows.NewWorkflowDefinitionHandler(&workflows.WorkflowCommandHandler{}, mockQuery)
			// when
			handler.GetWorkflow(ctx)
			// then
			Expect(w.statusCode).To(Equal(http.StatusBadRequest))
			Expect(w.body).ToNot(Equal(expectedWorkflow))
		})

		It("should error when retreiving workflow", func() {
			// given
			w := NewCustomResponseWriter()
			ctx, _ := gin.CreateTestContext(w)
			groupId := "testGroup01"
			workflowID := "testW01"
			ctx.Set(workflows.GroupIdParam, groupId)
			ctx.Set(workflows.WorkflowIdParam, workflowID)
			workflow := workflows.Workflow{Name: "testW01"}
			expectedWorkflow, err := json.Marshal(workflow)
			if err != nil {
				fmt.Println(err)
				return
			}

			mockQuery := workflows.NewMockWorkflowsQuery(mockCtrl)
			mockQuery.EXPECT().GetWorkflow(groupId, workflowID).Times(1).Return(nil, fmt.Errorf("Erorr"))
			handler := workflows.NewWorkflowDefinitionHandler(&workflows.WorkflowCommandHandler{}, mockQuery)
			// when
			handler.GetWorkflow(ctx)
			// then
			Expect(w.statusCode).To(Equal(http.StatusInternalServerError))
			Expect(w.body).ToNot(Equal(expectedWorkflow))
		})

		It("should error when not found error when retreiving workflow", func() {
			// given
			w := NewCustomResponseWriter()
			ctx, _ := gin.CreateTestContext(w)
			groupId := "testGroup01"
			workflowID := "testW01"
			ctx.Set(workflows.GroupIdParam, groupId)
			ctx.Set(workflows.WorkflowIdParam, workflowID)
			workflow := workflows.Workflow{Name: "testW01"}
			expectedWorkflow, err := json.Marshal(workflow)
			if err != nil {
				fmt.Println(err)
				return
			}

			mockQuery := workflows.NewMockWorkflowsQuery(mockCtrl)
			mockQuery.EXPECT().GetWorkflow(groupId, workflowID).Times(1).Return(nil, errors.NotFoundError{Message: "Group not found"})
			handler := workflows.NewWorkflowDefinitionHandler(&workflows.WorkflowCommandHandler{}, mockQuery)
			// when
			handler.GetWorkflow(ctx)
			// then
			Expect(w.statusCode).To(Equal(http.StatusNotFound))
			Expect(w.body).ToNot(Equal(expectedWorkflow))
		})
	})
})
