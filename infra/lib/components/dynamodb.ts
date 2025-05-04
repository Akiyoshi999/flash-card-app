import { Construct } from "constructs";
import * as cdk from "aws-cdk-lib";
import * as ddb from "aws-cdk-lib/aws-dynamodb";
import { EnvProps } from "../type";

export class DynamodbConstruct extends Construct {
  public readonly table: ddb.Table;
  constructor(scope: Construct, id: string, props: EnvProps) {
    super(scope, id);

    this.table = new ddb.Table(this, `Table-${props.stackNameSuffix}`, {
      tableName: `FlashCardApp-${props.stackNameSuffix}-Table`,
      partitionKey: { name: "pk", type: ddb.AttributeType.STRING },
      sortKey: { name: "sk", type: ddb.AttributeType.STRING },
      removalPolicy: cdk.RemovalPolicy.DESTROY,
      billingMode: ddb.BillingMode.PAY_PER_REQUEST,
    });
  }
}
