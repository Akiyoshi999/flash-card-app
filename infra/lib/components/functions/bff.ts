import { GoFunction } from "@aws-cdk/aws-lambda-go-alpha";
import { Construct } from "constructs";
import * as path from "path";
import { EnvProps, baseDir } from "../../type";
import * as dynamodb from "aws-cdk-lib/aws-dynamodb";
import { RetentionDays } from "aws-cdk-lib/aws-logs";
interface BffFunctionProps extends EnvProps {
  ddb: dynamodb.Table;
}

export class BffFunction extends Construct {
  public readonly function: GoFunction;

  constructor(scope: Construct, id: string, props: BffFunctionProps) {
    super(scope, id);

    const { ddb } = props;

    this.function = new GoFunction(
      this,
      `BffFunction-${props.stackNameSuffix}`,
      {
        entry: path.join(baseDir, "backend/cmd/bff"),
        environment: {
          DYNAMODB_TABLE_NAME: ddb.tableName,
        },
        logRetention: RetentionDays.ONE_DAY,
      }
    );

    // Grant permissions to the function
    ddb.grantReadWriteData(this.function);
  }
}
