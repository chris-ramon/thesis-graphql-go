# Contents

- [1. Introduction](#1-introduction)
- [2. Problem Statement](#2-problem-statement)
  - [2.1 Context](#21-context)
  - [2.2 Analysis](#22-analysis)
  - [2.3 Research Statement](#23-research-statement)
  - [2.4 Theory Justification](#24-theory-justification)
  - [2.5 Objectives](#25-objectives)
    - [2.5.1 General](#251-general)
    - [2.5.2 Specific](#252-specifics)
- [3. Theory Framework](#3-theory-framework)
  - [3.1 International Previous Work](#31-international-previous-work)
  - [3.2 National Previous Work](#32-national-previous-work)
  - [3.3 Fundamental Theories](#33-fundamental-theories)
  - [3.4 Glossary](#34-glossary)
- [4. Development](#4-development)
  - [4.1 Methodology](#41-methodology)
  - [4.2 Analysis](#42-analysis)
  - [4.3 Introduction](#43-introduction)
  - [4.4 Tasks](#44-tasks)
  - [4.5 Priorization](#45-priorization)
  - [4.6 Planning](#46-planning)
  - [4.7 Scope](#47-scope)
  - [4.8 Requirements](#48-requirements)
  - [4.9 Design](#49-design)
  - [4.10 Architecture](#410-architeture)
  - [4.11 Development](#411-develpoment)
  - [4.12 Testing](#412-testing)
  - [4.13 Deployment](#413-deployment)
  - [4.14 Artifacts](#414-artifacts)
- [5. Results](#5-results)
- [6. Conclusions](#6-conclusions)
- [7. Recomendations](#7-recomendations)
- [8. References](#8-references)

## 1. Introduction
This thesis develops the implementation of the GraphQL standard in the Go programming language and validates it's compatibility.

GraphQL is a communication protocol between clients and servers, the capacities are delegated to the servers defined via schemas, is the principal interface for making different operations such as: queries, mutations and subscriptions, was created for communicating servers and clients that are user interface oriented.

The implementation has two main source of development:
- The GraphQL specification.
- The GraphQL implementation in the JavaScript programming language.

The main objective of the implementation is to have compatibility with the JavaScript implementation up-to the v0.6.0 version.

The implementation creates a GitHub proyect, which is adopted as a code version control system and software development life cycle tool.

The implementation decides for the Go programming language because it's modern and is constantly adopted by new projects, it focuses on cloud computing therefore leverages all the benefits of GraphQL.

The implementation compatibility is verified by various validation frameworks such as:
- Compatibility unit tests.
- Functional tests.
- User interface oriented tests.

The result of the implementation is a open source software library.

The success of the implementation is confirmed by identifying the open source projects and companies that depends on the library so they wire them into their own software.

## 2. Problem Statement

### 2.1 Context

The traditional strategy for building software is by using tools that allow the communication between clients and servers using traditional protocols such as: RPC[1], SOAP[2], REST[3].

These tools centralizes the functionaing on the server side, where the data is defined so it can be interacted with, these strategies of communication generates various problems which are solved on demand with different solutions.

Multiple endpoints are created in order to cover multiple use cases which creates complexity to manage because the end-to-end requests are on the client side.

These different endpoints causes over-fetching and under-fetching of data, the consequence is that it requires multiple requests in order to consolidate the information, because the server decided which data to send therefore the payloads are greater or lower than expected.

One of the principle consequences of having multiple endpoints is that it creates nested requests causing complexity named: n + 1 problem.

Ineficiency at the maintainability side, because when changes are made on the server side the impact is on cascade to the different clients which depend on the server therefore the clients are more error prone.

There are scalability problems on the server side causing more band-width usage.

Costs increases because the data are send from the servers requires more capacity in order to attend the demand so more provisoning and resources are required.

Development experience are limited because it requires multiple tools in order to interact with each of the protocols that uses the servers causing negative impact to integrate new ones.
Slow development experience because multiple tools are required to interact with the different protocols which creates difficulties at the moment of integrate new team members to the proyects and also on the learning curve side.

Complexity at the application client side because it requires multiples layers of abstraction to coordinate the re-consolidate the information so the code are difficult to maintain because there are multiple solutions needs in order to address the problems.

The traditional protocols causes multiples strategies to manage the APIs versioning which triggers as result complexity on the server side that causes limited extensibility.

### 2.2 Analysis

#### Main research questions:

Has the implementation reached the defined state in order to fulfill the GraphQL features that were defined by the version v0.6.0 ?

Has the compatibility results shown that the quality standards were accomplished to meet the defined validations ?

#### Description of research questions:
Aims to investigate the first part of the research focusing on the implementation life cycle of the GraphQL standard.

Aims to verify the compatibility of the implementation by leveraging multiple validation frameworks.

#### Secondary research questions:

Which other GraphQL implementations are ?

Which Go best practices were decided to use ?

Why was it decided to build the implementation as an open source project ?

How was the compatibility verification measured ?

Which companies and open source projects use the software library ?

Which framework of maintenance was set up for the software library ?

#### Description of the research questions:

Aims to describe what other implementations are in order to decide which best practices to follow and have them as reference for this implementation.

Aims to describe the decision of leveraging the software development strategy of open source so the future maintenance is ensured for keeping the project alive.

Aims to describe the different strategies to verify the compatibility of the implementation.

Aims to describe which companies and open source projects leverage all the benefits of the implementation in their own software.

Aims to describe the different tooling around the software library that was setup in order to build and maintain the implementation.

### 2.3 Research Statement

Hernández and Mendoza (2018): The justification of arguments are needed for responding the reasonings of the study development, the importance of the problem and who affects. 

Consequently the justifications of the studay are elaborated.

#### 2.4 Theory Justification

Hernández and Mendoza (2018): Argues that value is created when justification of a theory is made, the result can be used by new researchers and to create new concepts.

It was observed that at the time of the release of the GraphQL specification in the year 2015 the implementations were limited only to JavaScript[10].

The implementation result can be used by different other researchers because it was decided to maintain it as an open source project.

It does create new concepts because at the time the implementation was released there were no other Go implementations so it created a based and foundational reference for upcoming new implementations.

#### Applied Justification

Hernández and Mendoza (2018): Argues that the justification of the pratice creates value when solving one or more real problems through innovation and technologies for the quality of the processes.

The investigation is justified because it creates value through the library and the quality of the implementation is guarantee through multiple validations frameworks.

#### Social Justification

Hernández and Mendoza (2018) argues that the social justification impacts the society and for the end users of the investigation.

The investigation has social impact because it has benefits for open source and companies because the implementation end result is open source and the impact is at full transparency because all the development is centralized in GitHub and due to the license the cost of usage is zero, and cascades infinity improvements because many contributors supports the library creating changes that allows the library to be future proof that avoids error regressions and security fixes.


### 2.5 Objectives
### 2.5.1 General

GraphQL implementation and compatibility verification.

### 2.5.2 Specifics

- Investigate the GraphQL standard implementation life cycle.

- Investigate the compatibility verification leveraging varios frameworks.

- Investigate current state of alternative implementations.

- Investigate the decision making for opting-in for open-source model.

- Investigate which companies use the end result software library.

- Investigate the tooling around the project.

# 3 Theory Framework
## 3.1 International Previous Work

99designs/gqlgen (2016/10/12): Implements a schema based graphql which generates all the components on the server side at build time, the key value is that all the parts are described on the graphql file which contains all the definitions of a type system and from there via using the command line interface commands it generates all the Go code which are: An http/s server, GraphQL endpoint and unit tests, and related development tools such as GraphQL Playground.

The drawbacks being that the Go code is auto-generated, which limits the extensibility, since the code is locked-in and making changes are difficult because it requires core dependency from the upstream project.


graph-gophers/graphql-go (2016/10/12): Implements the GraphQL standard by having as an entry-points the schema as a file, and the types resolvers are defined via code, it is used a library, main advantage being the adoption, also the core is shared across other implementations therefore the benefits are increased by the contributions from the open-source community.

The drawbacks being that the schema is defined at text level which have their limitations in terms of extensibility because it requires extra tooling to maintain the schema source of true.


dosco/graphjin (2019/03/24): Implements the GraphQL standard, being main features the ability to generate the API from a schema that generates the SQL queries, so it is an end-to-end solution similar than a object-relational-mapping, it does support multiple databases, allows end users to just focus on the GraphQL schema side so the Go code is generated, using the API enables to perform queries, mutations and subscriptions.

The drawbacks being that the end-to-end solution locks-in the implementation which is hard to extend, therefore the dependency on having internal hooks exposed and requires strong dependency on the maintainers.

## 3.2 National Previous Work

(Karla Cecilia Reyes Burgos, 2023) Created a research titled: Web services with GraphQL. Which is a sistematic literature, that details the recopilation of diverse resources which investigates the GraphQL usage, and has as a principal aim answer the following questions:

- ¿Which areas the science proves greater interest in the development of web services using GraphQL?
- The medicine industry is the greater area of science interested that applies a GraphQL implementation.

- ¿Which countries had shown greater interest in the develpoment of web services using GraphQL?
- The country with greater interest in apply a GraphQL implementation is USA.

- ¿In which year we find the greater quantity of investigations with respect of the development of web services using GraphQL?
- The year which were more informatic solutions were 2019.

## 3.3 Fundamental Theories

## 3.4 Glossary

Rapid Application Development (RAD): A methodology that emphasizes and prioritizes rapid iterations of software development.

Some of the key aspects being:
Prototyping: Leverages testing ideas and collects feedback before creating the final software version.
User Feedback: Focuses on user feedback and continuous iteration of small working parts instead of a strict plan.
Fast turnaround: Benefits the turnaround of the development of the software to pivot making it appealing for developers that are into highly fast paced working environments.

Open-source Development Methodology: A methodology that focuses on collaboration across different numbers of contributors, it is decentralized and in openness development.

Some of the key aspects being:
Peer Review: Continue development of parts that are improved by the community so the risk of errors are reduced as much as possible.

Decentralized Contributors: Development as a decentralized iterations via different tools that the community can leverage for synchronizing all the development workflows without single point of control.

Public Codebase: The source of the project is at a public repository so can be accessed by any contributor to leverage all the benefits that are tied to the license, and the key feature is that security fixes can be addressed.

Rapid Prototyping: Creates the right environment for rapid prototyping of new features, which enables the participation of all the community via the tooling of the source of the code.

Collaboration: Encourages all types of participation which can be from different experiences so the community of contributors are able to participate and learn in different ways.

# 4. Development
## 4.1 Methodology

#### Implementation

Rapid Application Development (RAD): We used this methodology in order to implement rapidly iteratively the GraphQL Go implementation following the graphql-js reference:

Prototyping: The GraphQL Go implementation was created from very basic initial interfaces that started from the text source, parser, lexer, AST, collectors and resolver.

User Feedback: As we made progress we received constant feedback that is documented mainly in GitHub, we did not follow a strict plan.

Fast turnaround: We did quick iterations in the implementation because it enabled us to constantly release multiple versions so in order to do so the contributors worked on a fast paced working environment, and also to cover the need of a working version for different companies that leverage the initial versions.

Open-source development methodology: The development of the project was by leveraging the different characteristics of the open-source methodology to have multiple contributors adding value to the project, it is decentralized so the project can be keep forever due to the dependency on GitHub which guarantee durability and also the development was continuously iterated in the openness so any person could benefit from the project from learning, up-to proposing changes.

Peer Review: The project was constantly peer-review by different persons across the planet, and the key part of those besides proposing improvements is that they reported and sent security fixes to keep the project trustable and safe from vulnerabilities.

Decentralized Contributors: The project was iterated and leveraged GitHub as a central control for the development; it enabled multiple contributors from different parts of the planet.

Public Codebase: The project was publicly available since the early stages of development, it allows other contributors to propose security fixes as well as reported bug plus security issues, the license is MIT which matches the reference graphql-js and enables creation of multiple branches of similar projects plus forks with improvements as well as internal forks that require special characteristics that cover inner needs of private projects.

Rapid Prototyping: Since we leveraged GitHub as a central part of the workflow of the project, it enabled rapid iterations of including sets of changes from the graphql-js reference implementation versions.

Collaboration: it enabled different contributors from different parts that learned from the project and they were to add value.

#### Validation

Validation process to verify that the software library implemented is successfully compatible with the standard.

Risk Assessment: Proactive approach of identifying and mitigating potential risk of the software library.

Functionality Testing: Process to check that the software library matched the standard.

Validation Testing: Process to ensure that the final software library meets the requirements of the stakeholders.

User Acceptance Testing: Process to test the software library with end-users in terms of metrics that proves their engagement and satisfaction aligned with the specification, the following criterias are collected and compared:
- GitHub Stars.
- GitHub Issues (Open/Closed).
- GitHub Issues (Open/Closed, comments that are related to the specification).
- GitHub Pull Requests (open/closed).
- GitHub Forks.
- GitHub Repository License.

Design Verification and Reviews: Process to oversee the software library design or architecture.

References: [IEEE 1012](https://github.com/chris-ramon/thesis-graphql-go/blob/b2626cc75dcba11aef045e6a28769411f25f7d26/README.md#8-references), [ISO 13485](https://github.com/chris-ramon/thesis-graphql-go/blob/b2626cc75dcba11aef045e6a28769411f25f7d26/README.md#8-references), [ISO 15026](https://github.com/chris-ramon/thesis-graphql-go/blob/b2626cc75dcba11aef045e6a28769411f25f7d26/README.md#8-references).


## 4.2 Analysis

#### Implementation

It was decided to implement the GraphQL standard in Go following the methodologies listed above to have a working initial version and continue iterating from there:

Rapid application development (RAD): From the initial version there was multiple number of iterations that made the library stable:

Prototyping: We have earlier versions that have the implementation work as a prototype which helped us to validate and test main functionalities.

User Feedback: Since we used GitHub as a central main source, there was continued feedback from the community in order to incorporate those.

Fast turnaround: There were multiple pivots on the implementation of the main components with the aim of friendly APIs.

Open-source development methodology: We leveraged GitHub as source of the code, and we used multiple of those features to create the library as a open-source project:

Peer Review: The pull-requests that included new functionality had constantly reviewed to improve the quality and reduce the number of errors.

Decentralized Contributors: We had multiple contributors via GitHub since they were able to review the library and propose improvements and report bugs.

Public Codebase: The library was developed publicly therefore the codebase could be leveraged for different purposes.

Rapid Prototyping: The library was created as continuous iteration via multiple prototypes.

Collaboration: Since GitHub were used, collaboration is the central part of the development with constant collaboration of multiple contributors.

#### Validation

The following frameworks for compatibility validation were created:

- [Compatibility Unit Tests](https://github.com/graphql-go/compatibility-unit-tests): Validation compatibility library for comparing GraphQL implementations unit tests results.
- [Compatibility Standard Definitions](https://github.com/graphql-go/compatibility-standard-definitions): Validation compatibility library for comparing GraphQL's standard definitions against implementations.
- [Compatibility User Acceptance](https://github.com/graphql-go/compatibility-user-acceptance): Validation compatibility library for GraphQL implementation user acceptance.

It was decided to choose JavaScript because there are more JavaScript compilers available.

Risk Assessment: Compatibility validation of risks, done by wiring the three frameworks created to the continous integration pipeline and in case of regressions against not fullfilling the standard there is warnings and newer changes always being validated for keeping the implementation stable, the following frameworks guarantee risk assessment:
- Compatibility Unit Tests.
- Compatibility Standard Definitions.
- Compatibility User Acceptance.
  
Functionality Testing: Compatibility validation against the JavaScript official reference implementation unit tests via the following framework:
- Compatibility Unit Tests.

Validation Testing & Design Verification and Reviews: Compatibility validation against the official standard documentation, via the following framework:
- Compatibility Standard Definitions. 

User Acceptance Testing: Compatibility validation via collecting acceptance information from GitHub, via the following framework:
- Compatibility User Acceptance.


#### Differences

Standard definitions and unit tests libraries have some cross-cover of validations, the standard definitions analysis are the type system level which enables detailed differences that could be leverage for cross-analysis of implementation versions.


#### Related Dependencies Improvements

##### tomarrell/wrapcheck

Documentation Improvements:
GitHub Repository: https://github.com/tomarrell/wrapcheck
GitHub Pull Request: https://github.com/tomarrell/wrapcheck/pull/62


## 4.3 Introduction

The implementation was decided to be built because at the time the GraphQL standard was released in 2015 there were no implementations available in the Go programming language therefore the opportunity to create an open source software library was available.

One important highlight is that the standard was released at similar time together with the reference implementation graphql-js therefore those two artifacts were used to identify the tasks that are part of this section.

The tasks matches the main functionalities of the reference implementation graphql-js, so the main components were the central part of each iteration via GitHub so we could accomplish implementation at each version.

## 4.4 Tasks

Listing the tasks that covers the implementation and compatibility validation:

<table>
  <tbody>
    <tr>
      <td colspan=5 align="center"><b>Implementation</b></td>
    </tr>
    <tr>
      <th align="center">Priority</th>
      <th align="center">Task Name</th>
      <th align="center">Duration Interval Dates (Created At - Merged At)</th>
      <th align="center">Duration In Days</th>
      <th align="center">Contributors</th>
      <th align="center">Components</th>
      <th align="center">Description</th>
      <th align="center">Pull Requests</th>
      <th align="center">Version</th>
    </tr>
    <tr>
      <td>1</td>
      <td>Porting changes from graphql-js version 0.4.18.</td>
      <td>2016-03-07 09:07:29 +0000 UTC - 2016-05-30 01:52:47 +0000 UTC</td>
      <td>83</td>
      <td>
- https://github.com/sogko</br>- https://github.com/coveralls</br>- https://github.com/pspeter3</br>- https://github.com/chris-ramon</br>- https://github.com/jvatic
      </td>
      <td>Errors, Languages, Types, Execution, Validation.</td>
      <td>
<p>Port changes from graphql-js up to the version v0.4.18 which includes the following functionalities:</p>
<p>- Consolidates the extension definition outside the type definition.</p>
<p>- Makes operation name optional.</p>
<p>- Compliance with the int sizing based on the specification.</p>
<p>- Changes on the function signature of graphql.NewTypeInfo</p>
<p>- Enables the possibility of removing the experimental FieldDefFn.</p>
      </td>
      <td>https://github.com/graphql-go/graphql/pull/117</td>
      <td>0.4.18.</td>
    </tr>
    <tr>
      <td>2</td>
      <td>Porting changes from graphql-js version 0.5.0.</td>
      <td>2016-04-18 08:46:26 +0000 UTC - 2016-06-11 21:13:46 +0000 UTC</td>
      <td>54</td>
      <td>
- https://github.com/sogko</br>- https://github.com/switchtrue</br>- https://github.com/bsr203</br>- https://github.com/coveralls</br>- https://github.com/kmulvey
      </td>      
      <td>Definition, Directives, Executor, Introspection, Rules, Schema, Types.</td>
      <td>
<p>Port changes from graphql-js up to the version v0.5.0 which includes the following functionalities:</p>
<p>- Improvements on introspection related to directive locations.</p>
<p>- Improvements in the schema language related to directives.</p>
<p>- Consolidates the `getTypeOf` method into the executor component.</p>
<p>- Schema changes related to types.</p>
<p>- Consolidate arguments including context to executor.</p>
<p>- Improvements in types overlapping in rules component.</p>
<p>- Add schema definition into language component.</p>
      </td>
      <td>https://github.com/graphql-go/graphql/pull/123</td>
      <td>0.5.0.</td>
    </tr>
    <tr>
      <td>3</td>
      <td>Porting changes from graphql-js version 0.6.0.</td>
      <td>2017-03-15 03:38:12 +0000 UTC - 2017-03-15 03:40:57 +0000 UTC</td>
      <td>1</td>
      <td>
- https://github.com/sogko</br>- https://github.com/coveralls</br>- https://github.com/F21</br>- https://github.com/tsunammis</br>- https://github.com/mishudark</br>- https://github.com/talk-to-my-car</br>- https://github.com/chris-ramon</br>- https://github.com/chris-ramon</br>- https://github.com/coveralls</br>- https://github.com/sogko"
      </td>
      <td>Definition, Directives, Executor, Introspection, Rules, Schema, Validator.</td>
      <td>
<p>Port changes from graphql-js up to the version v0.6.0 which includes the following functionalities:</p>
<p>- Deepen introspection query from 3 levels to 7.</p>
<p>- Improves validation error message when field names conflict.</p>
<p>- Improves validation error messages by listing suggestions.</p>
<p>- Fixes a bug where an empty "block" list could be skipped by the printer.</p>
<p>- Exports introspection in public API.</p>
<p>- Schema Language Directives.</p>
<p>- Directive location: schema definition.</p>
<p>- Deprecated directive tag.</p>
<p>- Validation: improving overlapping fields quality.</p>        
      </td>
      <td>https://github.com/graphql-go/graphql/pull/192</td>
      <td>0.6.0.</td>
    </tr>
    <tr>
      <td>4</td>
      <td>Executor</td>
      <td>2015-09-16 04:23:35 +0000 UTC - 2015-09-25 23:13:25 +0000 UTC</td>
      <td>9</td>
      <td>
- https://github.com/sogko</br>- https://github.com/chris-ramon"
      </td>
      <td>Errors, Executor, Language, Types.</td>
      <td>
<p>- Executor implementation.</p>
<p>- Partial implementation of resolve fields.</p>
<p>- Implemented most of the types.</p>
<p>- Partial schema definition implementation</p>
<p>- Added starwars graphql tests.</p>
      </td>
      <td>https://github.com/graphql-go/graphql/pull/8</td>
      <td>0.4.18</td>
    </tr>
    <tr>
      <td>5</td>
      <td>Source</td>
      <td>2015-09-13 09:36:05 +0000 UTC - 2015-09-14 23:09:05 +0000 UTC</td>
      <td>1</td>
      <td>
- https://github.com/chris-ramon
      </td>
      <td>Source.</td>
      <td>
<p>- Adding source api implementation.</p>
      </td>
      <td>https://github.com/graphql-go/graphql/pull/5</td>
      <td>0.4.18</td>
    </tr>
    <tr>
      <td>6</td>
      <td>Visitor</td>
      <td>2015-09-23 02:42:31 +0000 UTC - 2015-09-25 23:13:54 +0000 UTC</td>
      <td>2</td>
      <td>
- https://github.com/sogko</br>- https://github.com/chris-ramon
      </td>
      <td>Visitor, Printer, Validator.</td>
      <td>
<p>- Visitor wired to Printer and Validator.</p>
<p>- Visit component implementation.</p>
<p>- AST implementation with enter and leave functions wiring.</p>
<p>- Visitor implementation related to actions.</p>
<p>- AST reducer of visitor component.</p>
<p>- Schema parser tests.</p>
      </td>
      <td>https://github.com/graphql-go/graphql/pull/10</td>
      <td>0.4.18</td>
    </tr>
    <tr>
      <td>7</td>
      <td>Printer</td>
      <td>2015-09-23 02:42:31 +0000 UTC - 2015-09-25 23:13:54 +0000 UTC</td>
      <td>2</td>
      <td>
- https://github.com/sogko</br>- https://github.com/chris-ramon
      </td>
      <td>Printer.</td>
      <td>
<p>- Printer implementation work related to ast node transformations.</p>
<p>- Printer and visitor changes related to AST traversal.</p>
<p>- Printer changes related to node to string transformations.</p>
<p>- Printer and parser changes to transform string to node.</p>
      </td>
      <td>https://github.com/graphql-go/graphql/pull/10</td>
      <td>0.4.18</td>
    </tr>
    <tr>
      <td>8</td>
      <td>Parser</td>
      <td>2015-09-10 03:14:07 +0000 UTC - 2015-09-10 18:02:02 +0000 UTC</td>
      <td>1</td>
      <td>
- https://github.com/sogko</br>- https://github.com/chris-ramon
      </td>
      <td>Parser, AST.</td>
      <td>
<p>- Parser unit tests.</p>
<p>- AST types definitions.</p>
<p>- AST files ordering.</p>
      </td>
      <td>https://github.com/graphql-go/graphql/pull/2</td>
      <td>0.4.18</td>
    </tr>
    <tr>
      <td>9</td>
      <td>Lexer</td>
      <td>2015-09-10 12:11:05 +0000 UTC - 2015-09-11 05:30:14 +0000 UTC</td>
      <td>1</td>
      <td>
- https://github.com/sogko</br>- https://github.com/chris-ramon
      </td>
      <td>Lexer, Types, Errors.</td>
      <td>
<p>- Types improvements.</p>
      </td>
      <td>https://github.com/graphql-go/graphql/pull/3</td>
      <td>0.4.18</td>
    </tr>
    <tr>
      <td>10</td>
      <td>AST</td>
      <td>2015-09-10 12:11:05 +0000 UTC - 2015-09-11 05:30:14 +0000 UTC</td>
      <td>1</td>
      <td>
- https://github.com/sogko</br>- https://github.com/chris-ramon
      </td>
      <td>AST.</td>
      <td>
<p>- AST structs and definitions.</p>
<p>- AST types re-estructuring.</p>
      </td>
      <td>https://github.com/graphql-go/graphql/pull/3</td>
      <td>0.4.18</td>
    </tr>
    <tr>
      <td>11</td>
      <td>Collector</td>
      <td>2015-09-16 04:23:35 +0000 UTC - 2015-09-25 23:13:25 +0000 UTC</td>
      <td>9</td>
      <td>
- https://github.com/sogko</br>- https://github.com/chris-ramon
      </td>
      <td>Collector, Executor, Types.</td>
      <td>
<p>- Collector fields implementation.</p>
<p>- Collector work related types, schema definition and scalars.</p>
<p>- Collector related unit tests.</p>        
      </td>
      <td>https://github.com/graphql-go/graphql/pull/8</td>
      <td>0.4.18</td>
    </tr>
    <tr>
      <td>12</td>
      <td>Resolver</td>
      <td>2018-03-01 09:36:38 +0000 UTC - 2018-03-08 21:51:05 +0000 UTC</td>
      <td>7</td>
      <td>
- https://github.com/dvic</br>- https://github.com/coveralls</br>- https://github.com/chris-ramon
      </td>
      <td>Resolver.</td>
      <td>
<p>- Adds field resolver interface support.</p>
      </td>
      <td>
<p>- https://github.com/graphql-go/graphql/pull/288</p>
      </td>
      <td>0.4.18</td>
    </tr>
<tr>
      <td>13</td>
      <td>Resolver</td>
      <td>2015-09-16 04:23:35 +0000 UTC - 2015-09-25 23:13:25 +0000 UTC</td>
      <td>9</td>
      <td>
- https://github.com/sogko</br>- https://github.com/chris-ramon
      </td>
      <td>Resolver.</td>
      <td>
<p>- Adds resolver component implementation.</p>
      </td>
      <td>
<p>- https://github.com/graphql-go/graphql/pull/8</p>
      </td>
      <td>0.4.18</td>
    </tr>  
    <tr>
      <td>14</td>
      <td>Types</td>
      <td>2015-09-26 15:21:47 +0000 UTC - 2015-10-02 01:05:37 +0000 UTC</td>
      <td>5</td>
      <td>
- https://github.com/sogko</br>- https://github.com/chris-ramon</br>- https://github.com/chris-ramon</br>- https://github.com/sogko
      </td>
      <td>Types.</td>
      <td>
<p>- Added introspection integration with types.</p>
<p>- Consolidated types definitions.</p>
      </td>
      <td>https://github.com/graphql-go/graphql/pull/12</td>
      <td>0.4.18</td>
    </tr>
    <tr>
      <td>15</td>
      <td>Errors</td>
      <td>2018-11-28 04:14:47 +0000 UTC - 2018-12-03 01:16:35 +0000 UTC</td>
      <td>4</td>
      <td>
- https://github.com/chris-ramon</br>- https://github.com/limoli</br>- https://github.com/coveralls</br>- https://github.com/racerxdl
      </td>
      <td>Errors.</td>
      <td>
<p>- Adds original error support.</p>        
      </td>
      <td>https://github.com/graphql-go/graphql/pull/423</td>
      <td>0.4.18</td>
    </tr>
    <tr>
      <td>16</td>
      <td>CircleCI</td>
      <td>2018-07-19 03:27:49 +0000 UTC - 2018-07-19 03:41:48 +0000 UTC</td>
      <td>1</td>
      <td>
- https://github.com/chris-ramon</br>- https://github.com/coveralls</br>- https://github.com/chris-ramon</br>- https://github.com/chris-ramon</br>- https://github.com/coveralls
      </td>
      <td>CI, CD.</td>
      <td>
<p>- Replaces travis continous integration in favor of circle continous integration.</p>
      </td>
      <td>https://github.com/graphql-go/graphql/pull/361</td>
      <td>0.6.0</td>
    </tr>
    <tr>
      <td colspan=5 align="center"><b>Compatibility Validation</b></td>      
    </tr>
    <tr>
      <td></td>
      <td>Architecture.</td>
      <td></td>
      <td></td>
      <td></td>
      <td>Compatibility Unit Tests.</td>
      <td>
<p>- Create the architecture project for the compatibility unit tests framework.</p>
      </td>
      <td>        
<p>- https://github.com/graphql-go/compatibility-unit-tests/pull/8</p>
<p>- https://github.com/graphql-go/compatibility-unit-tests/pull/4</p>
      </td>
      <td>0.6.0</td>
    </tr>    
    <tr>
      <td></td>
      <td>Scaffolding.</td>
      <td></td>
      <td></td>
      <td></td>
      <td>Compatibility Unit Tests.</td>
      <td>
<p>- Create the scaffolding project for the compatibility unit tests framework.</p>
      </td>
      <td>
<p>- https://github.com/graphql-go/compatibility-unit-tests/pull/1</p>
<p>- https://github.com/graphql-go/compatibility-unit-tests/pull/2</p>
<p>- https://github.com/graphql-go/compatibility-unit-tests/pull/7</p>
<p>- https://github.com/graphql-go/compatibility-unit-tests/pull/6</p>
      </td>
      <td>0.6.0</td>
    </tr>
    <tr>
      <td></td>
      <td>Implementation.</td>
      <td></td>
      <td></td>
      <td></td>
      <td>Compatibility Unit Tests.</td>
      <td>
<p>- Implement the compatibility unit tests framework.</p>
      </td>
      <td>
<p>- https://github.com/graphql-go/compatibility-unit-tests/pull/5</p>
<p>- https://github.com/graphql-go/compatibility-unit-tests/pull/3</p>
<p>- https://github.com/graphql-go/compatibility-unit-tests/pull/9</p>        
      </td>
      <td>0.6.0</td>
    </tr>
    <tr>
      <td></td>
      <td>Architecture.</td>
      <td></td>
      <td></td>
      <td></td>
      <td>Compatibility Standard Definitions.</td>
      <td>
<p>- Create the architecture of the project for the compatibility standard definitions framework.</p>
      </td>
      <td></td>
      <td>0.6.0</td>
    </tr>
    <tr>
      <td></td>
      <td>Scaffolding.</td>
      <td></td>
      <td></td>
      <td>Compatibility Standard Definitions.</td>
      <td>
<p>- Create the scaffolding project for the compatibility standard definitions framework.</p>
      </td>
      <td></td>
      <td>0.6.0</td>
    </tr>
    <tr>
      <td></td>
      <td>Implementation.</td>
      <td></td>
      <td>Compatibility Standard Definitions.</td>
      <td>
<p>- Implement the compatibility standard definitions framework.</p>
      </td>
      <td></td>
      <td>0.6.0</td>
    </tr>
    <tr>
      <td></td>
      <td>Architecture.</td>
      <td></td>
      <td></td>
      <td></td>
      <td>Compatibility Standard Definitions.</td>
      <td>
<p>- Create the architecture of the project for the compatibility user acceptance framework.</p>
      </td>
      <td></td>
      <td>0.6.0</td>
    </tr>
    <tr>
      <td></td>
      <td>Scaffolding.</td>
      <td></td>
      <td></td>
      <td></td>
      <td>Compatibility User Acceptance.</td>
      <td>
<p>- Create the scaffolding project for the compatibility user acceptance framework.</p>
      </td>
      <td></td>
      <td>0.6.0</td>
    </tr>
    <tr>
      <td></td>
      <td>Implementation.</td>
      <td></td>
      <td></td>
      <td>Compatibility User Acceptance.</td>
      <td>
<p>- Implement the compatibility user acceptance framework.</p>
      </td>
      <td></td>
      <td>0.6.0</td>
    </tr>    
  </tbody>
</table>

### 4.6 Planning

Tasks mapping are based on matching the unit tasks against the project thesis requirements, below we list the tasks names against their requirements plus the main implementation components.

Task Name
Requerimientos Funcionales
GraphQL Component
Pull Request










### 4.7 Scope

Research goals:
Identify the problems of the traditional client-server protocols.
Implement the GraphQL standard in Go.
Produce a Go open-source library.
Verify the compatibility against the GraphQL reference implementation.
Create the environment for preserving the end-result.

Goals clarifications:
Problems identified were opted-in to consider only the following protocols: 
Implementation limited up-to graphql-js version: 0.6.0.
Specific features were implemented starting from graphql-js version: 0.6.0.



Potential biases identification:
Implementation is committed to open-source.
Standard implementation graphql-js compatible commitment.

Limits of the research:
Limited up-to a graphql-js version 0.6.0 tests suites.
Limited to GraphQL schema programmatically defined.

Not included in this research:
Other type of schema definition strategies such as: Inference from text schema-first.


### 4.8 Requirements

Related to how the system behaves at an internal level.
How all different apis of gql behave

Requerimientos funcionales:
Source: Source-in.
Parser: Source to AST and Lexer.
Lexer: Tokenizes the source-in.
AST: The source as a tree.
Collector: Matches AST to resolvers.
Resolver: Produces end-result, meaning the source-out.

Related to the usability, ux of the end-user, at external level.
Requerimientos no funcionales:

Performance:
Usability:
Scalability:
Security:
Maintainability:
Compatibility:
Portability:

Related to the domain category of the type of software, specific to the domain/industry.
Requerimientos de dominio:
GraphQL Standard Spec:
Open-source Standards:
Go Standards:

### 4.9 Design

What are the internal design of the GraphQL standard, let’s add here diagrams of the workflow, also let’s add each component workflow, prop using uml something like that, also, let’s add GraphQL Playground, GraphQL GraphiQL.

### 4.10 Architecture

query/mutation/subscription as source.
Source: Source-in.
Parser: Source to AST and Lexer.
Lexer: Tokenizes the source-in.
AST: The source as a tree.
Collector: Matches AST to resolvers.
Resolver: Produces end-result, meaning the source-out.

State machine is a representation of a workflow of nodes that produce as an end-result a given state, the nodes that produce the final state can be partial, meaning the paths of the end-result are defined per set.

GraphQL have multiple components that can be represented as a state machine, for example the most important component that are abstracted as:
AST state machine, up-down direction.
Source partial state machine, because one way flow.


### 4.11 Development

How is the distribution of the task accomplished ? Let’s add here screenshots of the PRs, maybe a table, let’s break the task into a simpler list and here let’s add more detailed tasks.

Also we could add a taks mapping against the components and linking with the code components in Go down to graphql js components.

### 4.12 Testing

Also let’s add percentage of accomplishing for reaching the desire covering tests results of graphql-js.

Also we are matching the tasks names against the set of unit tests in order to make sure the compatibility against the graphql-js reference implementation.


Task Name
GraphQL JS Unit Test
GraphQL Go Unit Test
Pull Request

### 4.13 Deployment

TODO: Add screenshots of each tool, for main components of gql.

The open-source project was centralized in GitHub, where we leverage different features that enabled the openness of the library, such as:

Comments: 
Pull Requests:
Issues:
Stats:
Documentation:
Releases:
Tags:

Also we had a continuous integration pipeline that guarantees that new functionality was fully tested against all the test suite, we leverage the following both providers: TravisCI and CircleCI.

TravisCI was deprecated in favor of CircleCI due to the latter having more modern functionality that was better for the future proof of the project.

The following features of the continuous integration are in place:
Building:
Per Go version building:




To guarantee that the suite of unit tests matched a high level standard set at a percentage accepted by the open-source community we leveraged Coveralls, and the following features were used:




Besides that we also leverage the Go’s feature “go doc” in order to self document the library

# 4.14 Artifacts

Artifacts were created as part of this thesis, most of them are Go open-source libraries that tackle
specific purposes and detailed below:

#### Implementation
##### `graphql-go/graphql`

GitHub Repository: https://github.com/graphql-go/graphql

Description:
- Go open-source library that implements the `graphql-js` reference implementation.


#### Compatibility
##### `graphql-go/compatibility-base`

GitHub Repository: https://github.com/graphql-go/compatibility-base

Description:

Go open-source library that has the base reusable code used by other `graphql-go/*` libraries.

##### `graphql-go/compatibility-unit-tests`

GitHub Repository: https://github.com/graphql-go/compatibility-unit-tests

Description:

Go open-source library that compares unit test names between a GraphQL reference implementation and other GraphQL implementations.

##### `graphql-go/compatibility-standard-definitions`

GitHub Repository: https://github.com/graphql-go/compatibility-standard-definitions

Description:

Go open-source library that compares the GraphQL reference implementation type system with other GraphQL implementations.


##### `graphql-go/compatibility-user-acceptance`

GitHub Repository: https://github.com/graphql-go/compatibility-user-acceptance

Description:

Go open-source library that compares acceptance criteria between the GraphQL reference implementation and other GraphQL implementations.

##### `chris-ramon/design-doc-self-generator`

GitHub Repository: https://github.com/chris-ramon/design-doc-self-generator

Description:
- Go open-source library that obtains GitHub information.


# 5. Results

TBC.

# 6. Conclusions

TBC.


# 7. Recomendations
What are the recommendations for future related work that is strongly tie to this thesis ?

#### Validation
- Usability Testing: UX testing was not covered as part of this thesis, but further work in regards the API design could be done, for example the most high level APIs could be extracted from the reference implementation and compared against the implementation. Also each implementation could be tested as if the APIs are using the most accepted best practices from the community as if it cover the best DX.

- Performance Testing: Similar than the compatibility libraries that were created as part of this thesis, another library could created to compare the response time of queries, mutations and subscriptions of a set of operations so reference and implementations could be compared.

# 8. References

- Facebook, Inc. (2019). GraphQL standard reference: https://spec.graphql.org/October2021

- GraphQL Spec: https://github.com/graphql/graphql-spec

- Official Website: https://graphql-go.github.io/graphql-go.org

- GraphQL Queries as state machine: https://rmosolgo.github.io/ruby/graphql/2016/11/12/graphql-query-as-a-state-machine.html

- GraphQL Spec License: https://jointdevelopment.org

- Open Systems Interconnection: https://aws.amazon.com/what-is/osi-model

- REST: https://datatracker.ietf.org/doc/html/rfc7231

- SOAP: https://datatracker.ietf.org/doc/html/rfc4227

- RPC: https://datatracker.ietf.org/doc/html/rfc5531 https://datatracker.ietf.org/doc/html/rfc1050

- GraphQL Implementations Releases: https://youtu.be/783ccP__No8?t=1253

- IEEE 1012: https://store.accuristech.com/standards/ieee-1012-2016?product_id=1901416&utm_source=google&utm_medium=cpc&utm_campaign=Omni%20|%20Google%20Ads%20|%20ACC-E%20|%20Top%20|%20Top%20|%20Non-Brand%20(Various%20Standards)%20|%20global%20-%20Dynamic&utm_content=Dynamic%20IEEE%20Products&utm_ad=725894111433&utm_term=&matchtype=&device=c&GeoLoc=9186197&placement=&network=g&campaign_id=22035107454&adset_id=172916797152&ad_id=725894111433&utm_term=&utm_campaign=Omni+%7C+Google+Ads+%7C+ACC-E+%7C+Non-Brand+(Various+Standards)+%7C+US+-+Dynamic+Ads&utm_source=adwords&utm_medium=ppc&hsa_acc=4039732030&hsa_cam=22035107454&hsa_grp=172916797152&hsa_ad=725894111433&hsa_src=g&hsa_tgt=dsa-2387125951358&hsa_kw=&hsa_mt=&hsa_net=adwords&hsa_ver=3&gad_source=1&gclid=CjwKCAiA74G9BhAEEiwA8kNfpQjAYrRLZNavr7HLzKzfSnMfksE2HTVxvBuvffHbIC7jmtO1oP0r1RoC8SwQAvD_BwE

- ISO 13485: https://www.iso.org/obp/ui/en/#iso:std:iso:13485:ed-3:v1:en

- ISO 15026: https://www.iso.org/standard/73567.html
