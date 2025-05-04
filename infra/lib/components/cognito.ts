import { Construct } from "constructs";
import { EnvProps } from "../type";
import * as cognito from "aws-cdk-lib/aws-cognito";
import { RemovalPolicy } from "aws-cdk-lib";
export class CognitoConstruct extends Construct {
  public readonly userPool: cognito.UserPool;
  public readonly userPoolClient: cognito.UserPoolClient;

  constructor(scope: Construct, id: string, props: EnvProps) {
    super(scope, id);

    this.userPool = new cognito.UserPool(
      this,
      `UserPool-${props.stackNameSuffix}`,
      {
        userPoolName: `FlashCardApp-${props.stackNameSuffix}-UserPool`,
        selfSignUpEnabled: true,
        autoVerify: { email: true },
        signInAliases: { email: true },
        standardAttributes: {
          email: { required: true, mutable: true },
        },
        removalPolicy: RemovalPolicy.DESTROY,
      }
    );

    this.userPoolClient = this.userPool.addClient("UserPoolAddClient", {
      userPoolClientName: `FlashCardApp-${props.stackNameSuffix}-UserPoolClient`,
      authFlows: {
        adminUserPassword: true,
        userPassword: true,
        userSrp: true,
      },
      generateSecret: false,
      oAuth: {
        flows: {
          implicitCodeGrant: true,
          authorizationCodeGrant: true,
        },
      },
    });
    // const domain = this.userPool.addDomain("Domain", {
    //   cognitoDomain: {
    //     domainPrefix: `flash-card-app-${props.stackNameSuffix}`,
    //   },
    // });
  }
}
