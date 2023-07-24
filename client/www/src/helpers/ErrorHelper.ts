export class ErrorHelper {
  static readonly errors = new Map<string, string>([["invalid credentials, either email not exist or invalid password", "Invalid Credential"]]);

  private static isError(error: unknown): error is Error {
    return error instanceof Error;
  }

  static toHumanReadable<T = unknown>(error: T): Error {
    const somethingWentWrongError = new Error("Something went wrong", { cause: error });
    if (!this.isError(error) || !this.errors.has(error.message)) return somethingWentWrongError;
    const err = new Error(this.errors.get(error.message), { cause: error });
    err.stack = error.stack;
    err.name = error.name;
    return err;
  }
}
