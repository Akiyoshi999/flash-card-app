#!/usr/bin/env node
import * as cdk from "aws-cdk-lib";
import { InfraStack } from "../lib/infra-stack";
import * as dotenv from "dotenv";

dotenv.config({ path: "../.env" });
const envName = process.env.ENV || "";

const app = new cdk.App();
new InfraStack(app, `InfraStack-${envName}`, {
  stackNameSuffix: envName,
  tags: {
    Project: "FlashCardApp",
    Env: envName,
  },
});
