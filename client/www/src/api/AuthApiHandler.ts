import * as Endpoint from "../common/endpoints";
import { ContentType, HttpMethod } from "../common/enums";
import { User, UserCredential } from "../common/shared";

export interface IAuthApiHandler {
  login(credential: UserCredential): Promise<{ token: string; user: User }>;
}

export class AuthApiHandler implements IAuthApiHandler {
  async login(credential: UserCredential): Promise<{ token: string; user: User }> {
    const response = await fetch(Endpoint.LOGIN, {
      method: HttpMethod.POST,
      headers: { ["Content-Type"]: ContentType.Json },
      body: JSON.stringify(credential),
    });

    if (!response.ok) {
      throw await this.handleFailedResponse(response);
    }
    return (await response.json()) as { token: string; user: User };
  }

  async handleFailedResponse(response: Response) {
    const contentType = response.headers.get("Content-Type");

    if (!contentType) return AuthApiHandler.errSomethingWrong;

    switch (true) {
      case contentType.includes(ContentType.TextPlain):
        return new Error(await response.text());
    }
    return AuthApiHandler.errSomethingWrong;
  }

  private static errSomethingWrong = new Error("Something went wrong");
}
