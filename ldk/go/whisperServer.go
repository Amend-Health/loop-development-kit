package ldk

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/open-olive/loop-development-kit/ldk/go/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// WhisperServer is used by the controller plugin host to receive plugin initiated communication
type WhisperServer struct {
	Authority Authority
	Impl      WhisperService
}

// WhisperMarkdown is used by loops to create markdown whispers
func (m *WhisperServer) WhisperMarkdown(ctx context.Context, req *proto.WhisperMarkdownRequest) (*emptypb.Empty, error) {
	session := NewSessionFromProto(req.Session)
	if err := m.Authority.ValidateSession(session); err != nil {
		return nil, err
	}

	err := m.Impl.Markdown(ctx, &WhisperContentMarkdown{
		Icon:     req.Meta.Icon,
		Label:    req.Meta.Label,
		Markdown: req.Markdown,
		Style: Style{
			BackgroundColor: req.Meta.Style.BackgroundColor,
			HighlightColor:  req.Meta.Style.HighlightColor,
			PrimaryColor:    req.Meta.Style.PrimaryColor,
		},
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// WhisperConfirm is used by loops to create confirm whispers
func (m *WhisperServer) WhisperConfirm(ctx context.Context, req *proto.WhisperConfirmRequest) (*proto.WhisperConfirmResponse, error) {
	session := NewSessionFromProto(req.Session)
	if err := m.Authority.ValidateSession(session); err != nil {
		return nil, err
	}

	response, err := m.Impl.Confirm(ctx, &WhisperContentConfirm{
		Icon:     req.Meta.Icon,
		Label:    req.Meta.Label,
		Markdown: req.Markdown,
		Style: Style{
			BackgroundColor: req.Meta.Style.BackgroundColor,
			HighlightColor:  req.Meta.Style.HighlightColor,
			PrimaryColor:    req.Meta.Style.PrimaryColor,
		},
		ResolveLabel: req.ResolveLabel,
		RejectLabel:  req.RejectLabel,
	})
	if err != nil {
		return nil, err
	}

	return &proto.WhisperConfirmResponse{
		Response: response,
	}, nil
}

// WhisperForm is used by loops to create form whispers
func (m *WhisperServer) WhisperForm(req *proto.WhisperFormRequest, stream proto.Whisper_WhisperFormServer) error {
	session := NewSessionFromProto(req.Session)
	if err := m.Authority.ValidateSession(session); err != nil {
		return err
	}

	inputs := make(map[string]WhisperContentFormInput, len(req.Inputs))
	for key, inputProto := range req.Inputs {
		key := key
		switch inputContainer := inputProto.InputOneof.(type) {
		case *proto.WhisperFormInput_Checkbox_:
			inputs[key] = &WhisperContentFormInputCheckbox{
				Label:   inputContainer.Checkbox.Label,
				Tooltip: inputContainer.Checkbox.Tooltip,
				Value:   inputContainer.Checkbox.Value,
				Order:   inputContainer.Checkbox.Order,
				OnChange: func(value bool) {
					err := stream.Send(&proto.WhisperFormStreamResponse{
						WhisperFormResponseOneof: &proto.WhisperFormStreamResponse_Update{
							Update: &proto.WhisperFormUpdate{
								Key: key,
								Output: &proto.WhisperFormOutput{
									OutputOneof: &proto.WhisperFormOutput_Checkbox_{
										Checkbox: &proto.WhisperFormOutput_Checkbox{
											Value: value,
										},
									},
								},
							},
						},
					})
					// TODO: Fix this when we refactor to sidekick
					if err != nil {
						fmt.Println("ldk.WhisperServer.WhisperForm -> Checkbox onChange: error => ", err.Error())
					}
				},
			}
		case *proto.WhisperFormInput_Email_:
			inputs[key] = &WhisperContentFormInputEmail{
				Label:   inputContainer.Email.Label,
				Tooltip: inputContainer.Email.Tooltip,
				Value:   inputContainer.Email.Value,
				Order:   inputContainer.Email.Order,
				OnChange: func(value string) {
					err := stream.Send(&proto.WhisperFormStreamResponse{
						WhisperFormResponseOneof: &proto.WhisperFormStreamResponse_Update{
							Update: &proto.WhisperFormUpdate{
								Key: key,
								Output: &proto.WhisperFormOutput{
									OutputOneof: &proto.WhisperFormOutput_Email_{
										Email: &proto.WhisperFormOutput_Email{
											Value: value,
										},
									},
								},
							},
						},
					})
					// TODO: Fix this when we refactor to sidekick
					if err != nil {
						fmt.Println("ldk.WhisperServer.WhisperForm -> Email onChange: error => ", err.Error())
					}
				},
			}
		case *proto.WhisperFormInput_Markdown_:
			inputs[key] = &WhisperContentFormInputMarkdown{
				Label:   inputContainer.Markdown.Label,
				Tooltip: inputContainer.Markdown.Tooltip,
				Value:   inputContainer.Markdown.Value,
				Order:   inputContainer.Markdown.Order,
				OnChange: func(value string) {
					err := stream.Send(&proto.WhisperFormStreamResponse{
						WhisperFormResponseOneof: &proto.WhisperFormStreamResponse_Update{
							Update: &proto.WhisperFormUpdate{
								Key: key,
								Output: &proto.WhisperFormOutput{
									OutputOneof: &proto.WhisperFormOutput_Markdown_{
										Markdown: &proto.WhisperFormOutput_Markdown{
											Value: value,
										},
									},
								},
							},
						},
					})
					// TODO: Fix this when we refactor to sidekick
					if err != nil {
						fmt.Println("ldk.WhisperServer.WhisperForm -> Markdown onChange: error => ", err.Error())
					}
				},
			}
		case *proto.WhisperFormInput_Number_:
			inputs[key] = &WhisperContentFormInputNumber{
				Label:   inputContainer.Number.Label,
				Tooltip: inputContainer.Number.Tooltip,
				Value:   inputContainer.Number.Value,
				Order:   inputContainer.Number.Order,
				OnChange: func(value float32) {
					err := stream.Send(&proto.WhisperFormStreamResponse{
						WhisperFormResponseOneof: &proto.WhisperFormStreamResponse_Update{
							Update: &proto.WhisperFormUpdate{
								Key: key,
								Output: &proto.WhisperFormOutput{
									OutputOneof: &proto.WhisperFormOutput_Number_{
										Number: &proto.WhisperFormOutput_Number{
											Value: value,
										},
									},
								},
							},
						},
					})
					// TODO: Fix this when we refactor to sidekick
					if err != nil {
						fmt.Println("ldk.WhisperServer.WhisperForm -> Number onChange: error => ", err.Error())
					}
				},
			}
		case *proto.WhisperFormInput_Password_:
			inputs[key] = &WhisperContentFormInputPassword{
				Label:   inputContainer.Password.Label,
				Tooltip: inputContainer.Password.Tooltip,
				Order:   inputContainer.Password.Order,
				OnChange: func(value string) {
					err := stream.Send(&proto.WhisperFormStreamResponse{
						WhisperFormResponseOneof: &proto.WhisperFormStreamResponse_Update{
							Update: &proto.WhisperFormUpdate{
								Key: key,
								Output: &proto.WhisperFormOutput{
									OutputOneof: &proto.WhisperFormOutput_Password_{
										Password: &proto.WhisperFormOutput_Password{
											Value: value,
										},
									},
								},
							},
						},
					})
					// TODO: Fix this when we refactor to sidekick
					if err != nil {
						fmt.Println("ldk.WhisperServer.WhisperForm -> Password onChange: error => ", err.Error())
					}
				},
			}
		case *proto.WhisperFormInput_Radio_:
			inputs[key] = &WhisperContentFormInputRadio{
				Label:   inputContainer.Radio.Label,
				Tooltip: inputContainer.Radio.Tooltip,
				Options: inputContainer.Radio.Options,
				Order:   inputContainer.Radio.Order,
				OnChange: func(value string) {
					err := stream.Send(&proto.WhisperFormStreamResponse{
						WhisperFormResponseOneof: &proto.WhisperFormStreamResponse_Update{
							Update: &proto.WhisperFormUpdate{
								Key: key,
								Output: &proto.WhisperFormOutput{
									OutputOneof: &proto.WhisperFormOutput_Radio_{
										Radio: &proto.WhisperFormOutput_Radio{
											Value: value,
										},
									},
								},
							},
						},
					})
					// TODO: Fix this when we refactor to sidekick
					if err != nil {
						fmt.Println("ldk.WhisperServer.WhisperForm -> Radio onChange: error => ", err.Error())
					}
				},
			}
		case *proto.WhisperFormInput_Select_:
			inputs[key] = &WhisperContentFormInputSelect{
				Label:   inputContainer.Select.Label,
				Tooltip: inputContainer.Select.Tooltip,
				Options: inputContainer.Select.Options,
				Order:   inputContainer.Select.Order,
				OnChange: func(value string) {
					err := stream.Send(&proto.WhisperFormStreamResponse{
						WhisperFormResponseOneof: &proto.WhisperFormStreamResponse_Update{
							Update: &proto.WhisperFormUpdate{
								Key: key,
								Output: &proto.WhisperFormOutput{
									OutputOneof: &proto.WhisperFormOutput_Select_{
										Select: &proto.WhisperFormOutput_Select{
											Value: value,
										},
									},
								},
							},
						},
					})
					// TODO: Fix this when we refactor to sidekick
					if err != nil {
						fmt.Println("ldk.WhisperServer.WhisperForm -> Select onChange: error => ", err.Error())
					}
				},
			}
		case *proto.WhisperFormInput_Tel_:
			inputs[key] = &WhisperContentFormInputTel{
				Label:   inputContainer.Tel.Label,
				Tooltip: inputContainer.Tel.Tooltip,
				Order:   inputContainer.Tel.Order,
				Value:   inputContainer.Tel.Value,
				OnChange: func(value string) {
					err := stream.Send(&proto.WhisperFormStreamResponse{
						WhisperFormResponseOneof: &proto.WhisperFormStreamResponse_Update{
							Update: &proto.WhisperFormUpdate{
								Key: key,
								Output: &proto.WhisperFormOutput{
									OutputOneof: &proto.WhisperFormOutput_Tel_{
										Tel: &proto.WhisperFormOutput_Tel{
											Value: value,
										},
									},
								},
							},
						},
					})
					// TODO: Fix this when we refactor to sidekick
					if err != nil {
						fmt.Println("ldk.WhisperServer.WhisperForm -> Tel onChange: error => ", err.Error())
					}
				},
			}
		case *proto.WhisperFormInput_Text_:
			inputs[key] = &WhisperContentFormInputText{
				Label:   inputContainer.Text.Label,
				Tooltip: inputContainer.Text.Tooltip,
				Value:   inputContainer.Text.Value,
				Order:   inputContainer.Text.Order,
				OnChange: func(value string) {
					err := stream.Send(&proto.WhisperFormStreamResponse{
						WhisperFormResponseOneof: &proto.WhisperFormStreamResponse_Update{
							Update: &proto.WhisperFormUpdate{
								Key: key,
								Output: &proto.WhisperFormOutput{
									OutputOneof: &proto.WhisperFormOutput_Text_{
										Text: &proto.WhisperFormOutput_Text{
											Value: value,
										},
									},
								},
							},
						},
					})
					// TODO: Fix this when we refactor to sidekick
					if err != nil {
						fmt.Println("ldk.WhisperServer.WhisperForm -> Text onChange: error => ", err.Error())
					}
				},
			}
		case *proto.WhisperFormInput_Time_:
			value, err := ptypes.Timestamp(inputContainer.Time.Value)
			if err != nil {
				return err
			}

			inputs[key] = &WhisperContentFormInputTime{
				Label:   inputContainer.Time.Label,
				Tooltip: inputContainer.Time.Tooltip,
				Order:   inputContainer.Time.Order,
				Value:   value,
				OnChange: func(value time.Time) {
					valueProto, err := ptypes.TimestampProto(value)
					// TODO: Fix this when we refactor to sidekick
					if err != nil {
						fmt.Println("ldk.WhisperServer.WhisperForm -> Time onChange: timestamp conversion error => ", err.Error())
					}

					err = stream.Send(&proto.WhisperFormStreamResponse{
						WhisperFormResponseOneof: &proto.WhisperFormStreamResponse_Update{
							Update: &proto.WhisperFormUpdate{
								Key: key,
								Output: &proto.WhisperFormOutput{
									OutputOneof: &proto.WhisperFormOutput_Time_{
										Time: &proto.WhisperFormOutput_Time{
											Value: valueProto,
										},
									},
								},
							},
						},
					})
					// TODO: Fix this when we refactor to sidekick
					if err != nil {
						fmt.Println("ldk.WhisperServer.WhisperForm -> Email onChange: error => ", err.Error())
					}
				},
			}
		default:
			return fmt.Errorf("input had unexpected type %T", inputContainer)
		}
	}

	submitted, outputs, err := m.Impl.Form(stream.Context(), &WhisperContentForm{
		Icon:     req.Meta.Icon,
		Label:    req.Meta.Label,
		Markdown: req.Markdown,
		Style: Style{
			BackgroundColor: req.Meta.Style.BackgroundColor,
			HighlightColor:  req.Meta.Style.HighlightColor,
			PrimaryColor:    req.Meta.Style.PrimaryColor,
		}, Inputs: inputs,
		CancelLabel: req.CancelLabel,
		SubmitLabel: req.SubmitLabel,
	})
	// TODO: Fix this when we refactor to sidekick
	if err != nil {
		fmt.Println("ldk.WhisperServer.WhisperForm -> m.Impl.Form: error => ", err.Error())
	}

	outputsProto := make(map[string]*proto.WhisperFormOutput, len(outputs))
	for key, output := range outputs {
		outputProto, err := output.ToProto()
		if err != nil {
			return err
		}
		outputsProto[key] = outputProto
	}

	err = stream.Send(&proto.WhisperFormStreamResponse{
		WhisperFormResponseOneof: &proto.WhisperFormStreamResponse_Result{
			Result: &proto.WhisperFormResult{
				Submitted: submitted,
				Outputs:   outputsProto,
			},
		},
	})

	return err
}
