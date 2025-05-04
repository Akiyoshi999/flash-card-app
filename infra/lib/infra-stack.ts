import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { DynamodbConstruct } from "./components/dynamodb";
import { AppsyncConstruct } from "./components/appsync";
import { CognitoConstruct } from "./components/cognito";

interface InfraStackProps extends cdk.StackProps {
  stackNameSuffix: string;
}

export class InfraStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props: InfraStackProps) {
    super(scope, id, props);

    const stackNameSuffix = props.stackNameSuffix;

    const cognito = new CognitoConstruct(this, `Cognito-${stackNameSuffix}`, {
      stackNameSuffix,
    });

    const ddb = new DynamodbConstruct(this, `DynamoDB-${stackNameSuffix}`, {
      stackNameSuffix,
    });

    const appsync = new AppsyncConstruct(this, `Appsync-${stackNameSuffix}`, {
      stackNameSuffix,
      table: ddb.table,
      userPool: cognito.userPool,
    });
  }
}
