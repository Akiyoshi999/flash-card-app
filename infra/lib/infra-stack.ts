import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { DynamodbConstruct } from "./components/dynamodb";
// import * as sqs from 'aws-cdk-lib/aws-sqs';

interface InfraStackProps extends cdk.StackProps {
  stackNameSuffix: string;
}

export class InfraStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props: InfraStackProps) {
    super(scope, id, props);

    const stackNameSuffix = props.stackNameSuffix;

    const ddb = new DynamodbConstruct(this, `DynamoDB-${stackNameSuffix}`, {
      stackNameSuffix,
    });
  }
}
