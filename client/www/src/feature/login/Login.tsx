import { Button, Col, Form, Input, Layout, Row, Typography } from "antd";
import { shallowEqual, useDispatch, useSelector } from "react-redux";
import { Link, useNavigate } from "react-router-dom";
import { authAction } from "../../action/user.action";
import { User, UserCredential } from "../../common/shared";
import { AppDispatch, AppRootState } from "../../store";
import "./Login.scss";

export const Login = () => {
  const [form] = Form.useForm<UserCredential>();
  const dispatch = useDispatch<AppDispatch>();
  const navigate = useNavigate();
  const user = useSelector<AppRootState, User | null>(({ user }) => user.info, shallowEqual);

  const onFormSubmitHandler = (formValue: UserCredential) => {
    void dispatch(authAction.login(formValue));
  };

  // useEffect(() => {
  //   if (user === null) return;
  // }, [user, navigate]);

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
