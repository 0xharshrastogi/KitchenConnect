import { Button, Col, Form, Input, Layout, Row, Typography, message } from "antd";
import { useDispatch } from "react-redux";
import { Link } from "react-router-dom";
import * as endpoint from "../../common/endpoints";
import { ContentType, HttpMethod } from "../../common/enums";
import { User, UserCredential } from "../../common/shared";
import { Func } from "../../common/types";
import { userActions } from "../../reducers";
import "./Login.scss";

type LoginApiResp = {
  token: string;
  user: User;
};

type FormSubmitResult<T> = { kind: "error"; message: string } | { kind: "success"; message: string; data: T };

const handleFailedResponse = async (response: Response): Promise<Error> => {
  switch (response.headers.get("Content-Type")) {
    case ContentType.TextPlain:
      return new Error(await response.text());
  }
  return new Error("Something went wrong");
};

const loginFormSubmitHandler: Func<[UserCredential], Promise<FormSubmitResult<LoginApiResp>>> = async (credential: UserCredential) => {
  const response = await fetch(endpoint.LOGIN, {
    method: HttpMethod.POST,
    headers: { "Content-Type": ContentType.Json },
    body: JSON.stringify(credential),
  });

  if (!response.ok) {
    const { message } = await handleFailedResponse(response);
    return { kind: "error", message };
  }

  if (response.headers.get("Content-Type")?.includes(ContentType.Json)) {
    const data = (await response.json()) as LoginApiResp;
    return { kind: "success", message: `Welcome back, ${data.user.firstName}`, data };
  }
  return { kind: "error", message: "something went wrong" };
};

export const Login = () => {
  const [form] = Form.useForm<UserCredential>();
  const dispatch = useDispatch();

  const onFormSubmitHandler = async (formValue: UserCredential) => {
    dispatch(userActions.reset({ loading: true }));

    const result = await loginFormSubmitHandler(formValue);
    const { kind, message: text } = result;

    if (kind === "error") {
      void message.error(text);
      return;
    }

    const { data } = result;
    void message.success(result.message);
    dispatch(userActions.activate(data.user));
  };

  return (
    <Row>
      <Col span={24} className="login-mobile">
        <Layout style={{ height: "100dvh" }}>
          <Layout.Header className="header login-mobile-content">
            <Typography.Title color="white">Sign up</Typography.Title>
          </Layout.Header>
          <Layout.Content className="content login-mobile-content">
            <div className="content-intro">
              <Typography.Title level={3}>Welcome Back</Typography.Title>
              <Typography.Text type="secondary">Hello there, sign in to continue</Typography.Text>
            </div>

            <Form
              autoComplete="off"
              className="form"
              form={form}
              onFinish={(value) => void onFormSubmitHandler(value)}
              labelCol={{ span: 8 }}
              wrapperCol={{ span: 16 }}
            >
              <fieldset>
                <Form.Item label="Email" name="email" rules={[{ required: true, message: "Please input your email!" }]}>
                  <Input className="form-control" type="email" placeholder="Enter your email" />
                </Form.Item>

                <Form.Item label="Password" name="password" rules={[{ required: true, message: "Please input your password!" }]}>
                  <Input.Password className="form-control" placeholder="Enter your password" />
                </Form.Item>
              </fieldset>

              <footer>
                <Button htmlType="submit" block type="primary">
                  Login
                </Button>
              </footer>
            </Form>
          </Layout.Content>
          <Layout.Footer style={{ textAlign: "center" }}>
            <footer>
              <Typography.Text>
                Don't have an account? <Link to="/signup">Sign up</Link>
              </Typography.Text>
            </footer>
          </Layout.Footer>
        </Layout>
      </Col>
    </Row>
  );
};
