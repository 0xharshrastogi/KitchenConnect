import { Button, Col, Form, Input, Layout, Row, Typography, message } from "antd";
import { Link, useNavigate } from "react-router-dom";
import { UserCredential } from "../../common/shared";
import { routes } from "../../constant/path";
import { useAuthentication } from "../../hooks";
import "./Login.scss";

export const Login = () => {
  const [form] = Form.useForm<UserCredential>();
  const auth = useAuthentication();
  const navigate = useNavigate();

  const onFormSubmitHandler = (formValue: UserCredential) => {
    void auth.login(formValue, {
      success: () => {
        void message.success("login success");
        navigate(routes.BASE);
      },
      error: (error) => void message.error(error.message),
    });
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
                  <Input.Password aria-required="true" autoComplete="true" className="form-control" placeholder="Enter your password" />
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
