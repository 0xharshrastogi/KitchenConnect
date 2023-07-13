import {
  Button,
  Col,
  Form,
  Input,
  Layout,
  Row,
  Typography,
  message,
} from "antd";
import { Link } from "react-router-dom";
import * as endpoint from "../../common/endpoints";
import { ContentType, HttpMethod } from "../../common/enums";
import { User, UserCredential } from "../../common/shared";
import "./Login.scss";

type TLoginResponse = {
  token: string;
  user: User;
};

export const Login = () => {
  const [form] = Form.useForm<UserCredential>();

  const onFormSubmit = async (credential: UserCredential) => {
    const response = await fetch(endpoint.LOGIN, {
      method: HttpMethod.POST,
      headers: {
        "Content-Type": ContentType.Json,
      },
      body: JSON.stringify(credential),
    });

    if (
      !response.ok &&
      response.headers.get("Content-Type")?.includes(ContentType.TextPlain)
    ) {
      void message.error("Invalid credentials", 2);
      return;
    }

    if (
      response.ok &&
      response.headers.get("Content-Type")?.includes(ContentType.Json)
    ) {
      const { user } = (await response.json()) as TLoginResponse;
      void message.success(`Welcome back, ${user.firstName}`, 2);
    }
  };

  const onFormSubmitHandler = (formValue: UserCredential) => {
    void onFormSubmit(formValue);
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
              <Typography.Text type="secondary">
                Hello there, sign in to continue
              </Typography.Text>
            </div>

            <Form
              autoComplete="off"
              className="form"
              form={form}
              onFinish={onFormSubmitHandler}
              labelCol={{ span: 8 }}
              wrapperCol={{ span: 16 }}
            >
              <fieldset>
                <Form.Item
                  label="Email"
                  name="email"
                  rules={[
                    { required: true, message: "Please input your email!" },
                  ]}
                >
                  <Input
                    className="form-control"
                    type="email"
                    placeholder="Enter your email"
                  />
                </Form.Item>

                <Form.Item
                  label="Password"
                  name="password"
                  rules={[
                    { required: true, message: "Please input your password!" },
                  ]}
                >
                  <Input.Password
                    className="form-control"
                    placeholder="Enter your password"
                  />
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
