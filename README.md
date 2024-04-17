# IMPORTANT
This is an experimental copy and not a fork of the replicated SDK located at [https://github.com/replicatedhq/replicated-sdk](https://github.com/replicatedhq/replicated-sdk).

**This is purposely not a fork**

## Purpose
The purpose of this test repo is to experiment with various types of SLSA based tools to create provenance and attestations. 

If you were invited to join the test, then you'll know that this repo is meant to experiment with using the SLSA Github Generator for Containers. 

See the [SLSA Github Generator for Containers](https://github.com/slsa-framework/slsa-github-generator/blob/main/internal/builders/container/README.md) repo.


## Project Scope

1. Using the `publish-with-slsa.yml` workflow, maintain the existing image build process and add provenance and attestations using the [SLSA Github Generator for Containers](https://github.com/slsa-framework/slsa-github-generator/blob/main/internal/builders/container/README.md) workflow.


### Requirements
1. Do not modify the `build-push-action` action, unless you find it absolutely necessary to meet the project objective. 
2. Do not copy, create, alter, make a local copy, etc of the `slsa-framework/slsa-github-generator` workflow, unless you find it absolutely necessary to meet the project objective. The purpose of the test is to try and utilize the SLSA Github Generator for Containers in its current state.
3. You must continue to build the image using the melange and apko Chainguard actions as defined in the `build-push-action`

## Notes
1. All the automated test suites have been commented out in the `publish-with-slsa` for ease of development. You don't need to add them back in.
2. You will need to supply your own docker hub credentials. We recommend forking this repo and adding your own credentials as repo secrets in your own fork.



## Contributing
The original contributing docs of the Replicated SDK remains in effect. Any code contributed here may or may not be submitted to the original repo. The Replicated SDK is a public repo and open source. If you would like to contribute to the Replicated SDK, then please do so using the official repo.


## License
The original license of the Replicated SDK remains in effect. Any code here is considered to be in the same license of the original repo.



------------
The original README from the actual repo is below:

# Introduction

The Replicated SDK (software development kit) is a service that allows you to embed key Replicated features alongside your application. 

[Replicated SDK Product Documentation](https://docs.replicated.com/vendor/replicated-sdk-overview) 

