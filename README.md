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
  - [4.10 Architecture](#410-architecture)
  - [4.11 Development](#411-development)
  - [4.12 Testing](#412-testing)
  - [4.13 Deployment](#413-deployment)
  - [4.14 Artifacts](#414-artifacts)
- [5. Results](#5-results)
- [6. Conclusions](#6-conclusions)
- [7. Recomendations](#7-recomendations)
- [8. References](#8-references)

## 1. Introduction
This thesis developed the implementation of the GraphQL standard in the Go programming language and validated its compatibility.

GraphQL is a communication protocol between clients and servers, where capacities are delegated to servers defined via schemas. It is the principal interface for making different operations such as: queries, mutations and subscriptions, and was created for communicating between servers and clients that are user interface oriented.

The implementation had two main sources of development:
- The GraphQL specification.
- The GraphQL implementation in the JavaScript programming language.

The main objective of the implementation was to have compatibility with the JavaScript implementation up to the v0.6.0 version.

The implementation created a GitHub project, which was adopted as a code version control system and software development life cycle tool.

The implementation chose the Go programming language because it is modern and is constantly adopted by new projects, it focuses on cloud computing therefore leverages all the benefits of GraphQL.

The implementation compatibility was verified by various validation frameworks such as:
- Compatibility unit tests.
- Functional tests.
- User interface oriented tests.

The result of the implementation is an open source software library.

The success of the implementation was confirmed by identifying the open source projects and companies that depend on the library and wire it into their own software.

## 2. Problem Statement

### 2.1 Context

The traditional strategy for building software is by using tools that allow communication between clients and servers using traditional protocols such as: RPC[1], SOAP[2], REST[3].

These tools centralize the functioning on the server side, where the data is defined so it can be interacted with. These strategies of communication generate various problems which are solved on demand with different solutions.

Multiple endpoints are created in order to cover multiple use cases which creates complexity to manage because the end-to-end requests are on the client side.

These different endpoints cause over-fetching and under-fetching of data. The consequence is that it requires multiple requests in order to consolidate the information, because the server decides which data to send therefore the payloads are greater or lower than expected.

One of the principal consequences of having multiple endpoints is that it creates nested requests causing complexity named: n + 1 problem.

Inefficiency at the maintainability side occurs because when changes are made on the server side the impact cascades to the different clients which depend on the server therefore the clients are more error prone.

There are scalability problems on the server side causing more bandwidth usage.

Costs increase because the data sent from the servers requires more capacity in order to attend the demand so more provisioning and resources are required.

Development experience is limited because it requires multiple tools in order to interact with each of the protocols that the servers use, causing negative impact when integrating new ones.
Slow development experience occurs because multiple tools are required to interact with the different protocols which creates difficulties at the moment of integrating new team members to the projects and also on the learning curve side.

Complexity at the application client side exists because it requires multiple layers of abstraction to coordinate and re-consolidate the information so the code is difficult to maintain because there are multiple solutions needed in order to address the problems.

The traditional protocols cause multiple strategies to manage the APIs versioning which triggers complexity on the server side that causes limited extensibility.

### 2.2 Analysis

#### Main research questions:

Did the implementation reach the defined state in order to fulfill the GraphQL features that were defined by the version v0.6.0?

Have the compatibility results shown that the quality standards were accomplished to meet the defined validations?

#### Description of research questions:
The first question aims to investigate the implementation life cycle of the GraphQL standard.

The second question aims to verify the compatibility of the implementation by leveraging multiple validation frameworks.

#### Secondary research questions:

Which other GraphQL implementations exist?

Which Go best practices were decided to use?

Why was it decided to build the implementation as an open source project?

How was the compatibility verification measured?

Which companies and open source projects use the software library?

Which framework of maintenance was set up for the software library?

#### Description of the research questions:

This aims to describe what other implementations exist in order to decide which best practices to follow and have them as reference for this implementation.

This aims to describe the decision of leveraging the software development strategy of open source so the future maintenance is ensured for keeping the project alive.

This aims to describe the different strategies to verify the compatibility of the implementation.

This aims to describe which companies and open source projects leverage all the benefits of the implementation in their own software.

This aims to describe the different tooling around the software library that was set up in order to build and maintain the implementation.

### 2.3 Research Statement

Hernández and Mendoza (2018): The justification of arguments is needed for responding to the reasonings of the study development, the importance of the problem and who it affects. 

Consequently the justifications of the study are elaborated.

#### 2.4 Theory Justification

Hernández and Mendoza (2018): Argue that value is created when justification of a theory is made, the result can be used by new researchers and to create new concepts.

It was observed that at the time of the release of the GraphQL specification in the year 2015 the implementations were limited only to JavaScript[10].

The implementation result can be used by other researchers because it was decided to maintain it as an open source project.

It does create new concepts because at the time the implementation was released there were no other Go implementations so it created a base and foundational reference for upcoming new implementations.

#### Applied Justification

Hernández and Mendoza (2018): Argue that the justification of the practice creates value when solving one or more real problems through innovation and technologies for the quality of the processes.

The investigation is justified because it creates value through the library and the quality of the implementation is guaranteed through multiple validation frameworks.

#### Social Justification

Hernández and Mendoza (2018) argue that the social justification impacts the society and the end users of the investigation.

The investigation has social impact because it has benefits for open source and companies because the implementation end result is open source and the impact is at full transparency because all the development is centralized in GitHub and due to the license the cost of usage is zero, and cascades infinite improvements because many contributors support the library creating changes that allow the library to be future proof that avoids error regressions and security fixes.


### 2.5 Objectives
### 2.5.1 General

GraphQL implementation and compatibility verification.

### 2.5.2 Specifics

- Investigate the GraphQL standard implementation life cycle.

- Investigate the compatibility verification leveraging various frameworks.

- Investigate current state of alternative implementations.

- Investigate the decision making for opting-in for open-source model.

- Investigate which companies use the end result software library.

- Investigate the tooling around the project.

# 3 Theory Framework
## 3.1 International Previous Work

99designs/gqlgen (2016/10/12): Implements a schema based GraphQL which generates all the components on the server side at build time. The key value is that all the parts are described in the GraphQL file which contains all the definitions of a type system and from there via using the command line interface commands it generates all the Go code which include: An HTTPS server, GraphQL endpoint and unit tests, and related development tools such as GraphQL Playground.

The drawbacks are that the Go code is auto-generated, which limits the extensibility, since the code is locked-in and making changes is difficult because it requires core dependency from the upstream project.


graph-gophers/graphql-go (2016/10/12): Implements the GraphQL standard by having as entry-points the schema as a file, and the types resolvers are defined via code. It is used as a library, with the main advantage being the adoption, also the core is shared across other implementations therefore the benefits are increased by the contributions from the open-source community.

The drawbacks are that the schema is defined at text level which has limitations in terms of extensibility because it requires extra tooling to maintain the schema source of truth.


dosco/graphjin (2019/03/24): Implements the GraphQL standard, with main features being the ability to generate the API from a schema that generates the SQL queries, so it is an end-to-end solution similar to an object-relational-mapping. It supports multiple databases, allows end users to just focus on the GraphQL schema side so the Go code is generated, using the API enables performing queries, mutations and subscriptions.

The drawbacks are that the end-to-end solution locks-in the implementation which is hard to extend, therefore the dependency on having internal hooks exposed and requires strong dependency on the maintainers.

## 3.2 National Previous Work

(Karla Cecilia Reyes Burgos, 2023) Created a research titled: Web services with GraphQL. This is a systematic literature that details the recompilation of diverse resources which investigate the GraphQL usage, and has as a principal aim to answer the following questions:

- Which areas of science show greater interest in the development of web services using GraphQL?
- The medicine industry is the greater area of science interested that applies a GraphQL implementation.

- Which countries had shown greater interest in the development of web services using GraphQL?
- The country with greater interest in applying a GraphQL implementation is USA.

- In which year do we find the greater quantity of investigations with respect to the development of web services using GraphQL?
- The year which had more informatic solutions was 2019.

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

Rapid Application Development (RAD): We used this methodology in order to implement the GraphQL Go implementation rapidly and iteratively following the graphql-js reference:

Prototyping: The GraphQL Go implementation was created from very basic initial interfaces that started from the text source, parser, lexer, AST, collectors and resolver.

User Feedback: As we made progress we received constant feedback that was documented mainly in GitHub; we did not follow a strict plan.

Fast turnaround: We did quick iterations in the implementation because it enabled us to constantly release multiple versions so in order to do so the contributors worked in a fast paced working environment, and also to cover the need of a working version for different companies that leveraged the initial versions.

Open-source development methodology: The development of the project was done by leveraging the different characteristics of the open-source methodology to have multiple contributors adding value to the project. It is decentralized so the project can be kept forever due to the dependency on GitHub which guarantees durability and also the development was continuously iterated in the open so any person could benefit from the project from learning, up to proposing changes.

Peer Review: The project was constantly peer-reviewed by different persons across the planet, and the key part of this besides proposing improvements was that they reported and sent security fixes to keep the project trustable and safe from vulnerabilities.

Decentralized Contributors: The project was iterated and leveraged GitHub as a central control for the development; it enabled multiple contributors from different parts of the planet.

Public Codebase: The project was publicly available since the early stages of development, it allowed other contributors to propose security fixes as well as report bugs and security issues. The license is MIT which matches the reference graphql-js and enables creation of multiple branches of similar projects plus forks with improvements as well as internal forks that require special characteristics that cover inner needs of private projects.

Rapid Prototyping: Since we leveraged GitHub as a central part of the workflow of the project, it enabled rapid iterations of including sets of changes from the graphql-js reference implementation versions.

Collaboration: It enabled different contributors from different parts that learned from the project and they were able to add value.

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

Risk Assessment: Compatibility validation of risks was done by wiring the three frameworks created to the continuous integration pipeline and in case of regressions against not fulfilling the standard there are warnings and newer changes are always being validated for keeping the implementation stable. The following frameworks guarantee risk assessment:
- Compatibility Unit Tests.
- Compatibility Standard Definitions.
- Compatibility User Acceptance.
  
Functionality Testing: Compatibility validation against the JavaScript official reference implementation unit tests was done via the following framework:
- Compatibility Unit Tests.

Validation Testing & Design Verification and Reviews: Compatibility validation against the official standard documentation was done via the following framework:
- Compatibility Standard Definitions. 

User Acceptance Testing: Compatibility validation via collecting acceptance information from GitHub was done via the following framework:
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

One important highlight is that the standard was released at a similar time together with the reference implementation graphql-js therefore those two artifacts were used to identify the tasks that are part of this section.

The tasks match the main functionalities of the reference implementation graphql-js, so the main components were the central part of each iteration via GitHub so we could accomplish implementation at each version.

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
      <td>12</td>
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
      <td>13</td>
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
      <td>14</td>
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
      <td>15</td>
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

#### Components

##### Implementation

The graphql-go library contains the following components:

- GraphQL:
 - Source.
 - Extensions.
- GQLErrors.
- Language:
  - AST.
  - Kinds.
  - Lexer.
  - Location.
  - Parser.
  - Printer.
  - Visitor.
- Executor:
  - Resolver.
- Validator.
- Definition.
- Directives.
- Introspection.
- Rules.
- Schema.

##### GraphQL
GraphQL is the component that wraps most of the internal components, it contains the end-user Go APIs, which are
syntactically similar to the graphql-js reference implementation APIs.

###### `graphql-js` Main API

```
import {
  graphql,
} from 'graphql';

graphql(schema, query).then(result => {
});
```

###### `graphql-go` Main API
```
import "github.com/graphql-go/graphql"

params := graphql.Params{Schema: schema, RequestString: query}
graphql.Do(params)
```

The internal components wrapped are the following:
- Source.
- Extensions.
- Parse.
- Validate.
- Execute.

##### Language Source
Source is the component that contains the GraphQL root operation in byte format.

Used within the Parse component as an entry point for accessing the library end-user GraphQL string operation.

##### Extensions
Extensions is a component that provides an interface for extending GraphQL execution with custom functionality through hooks into various phases of the GraphQL execution lifecycle.

The Extensions interface defines methods that allow developers to intercept and extend the following phases:
- **Init**: Initialize the extension with context and parameters
- **ParseDidStart**: Hook called before GraphQL query parsing begins
- **ValidationDidStart**: Hook called before GraphQL query validation begins  
- **ExecutionDidStart**: Hook called before GraphQL query execution begins
- **ResolveFieldDidStart**: Hook called before individual field resolution begins

Each phase hook returns a corresponding finish function that is called when the operation completes, allowing extensions to perform cleanup, logging, metrics collection, or other post-processing tasks. This design enables powerful extensibility for cross-cutting concerns like performance monitoring, authentication, caching, and custom validation logic.

##### GQLErrors

GQLErrors is the component responsible for modeling, collecting, and propagating errors that occur during GraphQL execution. It provides structures for representing errors with context such as locations, paths, and original error messages, aligning with the GraphQL specification’s error format. The component is based on the upstream [gqlerrors](https://github.com/graphql-go/graphql/tree/master/gqlerrors) package and ensures that errors are reported in a consistent and extensible way. This enables clients to receive detailed and actionable error information in GraphQL responses.

##### Language AST
AST (Abstract Syntax Tree) is the component that represents the parsed GraphQL document as a hierarchical tree structure of nodes.

The AST serves as the intermediate representation between the raw GraphQL query string and the execution engine. It is produced by the Parser component from the tokenized source and consumed by various other components including the Validator, Executor, and Printer.

The AST structure consists of different node types that correspond to GraphQL language constructs:
- **Document**: The root node representing the entire GraphQL document
- **OperationDefinition**: Represents query, mutation, or subscription operations
- **FieldDefinition**: Represents individual fields within operations
- **FragmentDefinition**: Represents reusable query fragments
- **SelectionSet**: Represents groups of field selections
- **Arguments**: Represents field arguments and their values
- **Directives**: Represents GraphQL directives applied to fields or fragments

The AST enables the GraphQL engine to traverse and analyze the query structure programmatically, allowing for validation rules to be applied, execution planning to be performed, and query introspection capabilities to be provided. This tree representation abstracts away the textual syntax and provides a structured format that can be efficiently processed by the subsequent stages of GraphQL execution.

##### Language Kinds
Kinds is the component that defines constants for identifying different types of AST nodes in the GraphQL document structure.

The Kinds component serves as a central registry of node type identifiers that are used throughout the GraphQL processing pipeline. It provides string constants that categorize each AST node according to its syntactic role in the GraphQL language.

The Kinds constants include:
- **Document**: Identifies the root document node
- **OperationDefinition**: Identifies query, mutation, and subscription operations
- **VariableDefinition**: Identifies variable declarations in operations
- **SelectionSet**: Identifies groups of field selections
- **Field**: Identifies individual field selections
- **Argument**: Identifies field arguments
- **FragmentSpread**: Identifies fragment usage
- **InlineFragment**: Identifies inline fragment selections
- **FragmentDefinition**: Identifies reusable fragment definitions
- **Variable**: Identifies variable references
- **IntValue**, **FloatValue**, **StringValue**, **BooleanValue**: Identify literal values
- **ListValue**, **ObjectValue**: Identify complex value structures
- **Directive**: Identifies GraphQL directives
- **NamedType**, **ListType**, **NonNullType**: Identify type references

The Kinds component enables type-safe AST traversal and manipulation by providing a standardized way to identify and categorize nodes. This is essential for the Visitor pattern implementation, validation rules, and other components that need to process specific types of AST nodes. The constants ensure consistency across the entire GraphQL processing pipeline and facilitate debugging by providing human-readable node type identifiers.

##### Language Lexer
Lexer is the component responsible for tokenizing the GraphQL source string into a sequence of meaningful tokens that can be processed by the Parser component.

The Lexer performs lexical analysis by scanning through the raw GraphQL query string character by character and identifying distinct tokens such as:
- **Names**: Field names, type names, and identifiers
- **Keywords**: GraphQL language keywords like `query`, `mutation`, `subscription`, `fragment`
- **Punctuation**: Braces `{}`, brackets `[]`, parentheses `()`, colons `:`, commas `,`
- **Operators**: Equality operators, directives `@`, spread operators `...`
- **Literals**: String literals, numeric literals, boolean literals
- **Comments**: Single-line and multi-line comments that are preserved or ignored as needed

The tokenization process transforms the linear text input into a structured sequence of tokens, each containing:
- **Token Type**: The category of the token (e.g., NAME, STRING, NUMBER)
- **Value**: The actual text content of the token
- **Location**: Position information including line and column numbers for error reporting

The Lexer serves as the foundation for the parsing pipeline, providing clean, categorized input to the Parser component which then constructs the AST. This separation of concerns allows for efficient error detection and reporting at the lexical level, helping developers identify syntax issues in their GraphQL queries before they reach the parsing stage.

##### Language Parser
Parser is the component responsible for analyzing the sequence of tokens produced by the Lexer and constructing an Abstract Syntax Tree (AST) that represents the hierarchical structure of a GraphQL document.

The Parser performs syntactic analysis by consuming tokens from the Lexer and applying GraphQL grammar rules to build a structured representation of the query. It processes the tokens according to the GraphQL specification and creates corresponding AST nodes for each language construct.

The Parser handles various GraphQL language elements including:
- **Document Parsing**: Processes the root level of a GraphQL document containing operations and fragments
- **Operation Parsing**: Handles query, mutation, and subscription operations with their selection sets
- **Selection Set Parsing**: Processes field selections, fragment spreads, and inline fragments
- **Field Parsing**: Handles field names, arguments, aliases, and nested selections
- **Fragment Parsing**: Processes both named fragment definitions and inline fragments
- **Value Parsing**: Handles literals (strings, numbers, booleans), variables, lists, and objects
- **Type Parsing**: Processes type references including named types, list types, and non-null types
- **Directive Parsing**: Handles directive applications with their arguments

The parsing process includes comprehensive error handling and reporting:
- **Syntax Error Detection**: Identifies malformed GraphQL syntax and provides precise error locations
- **Error Recovery**: Attempts to continue parsing after encountering errors to find additional issues
- **Location Tracking**: Maintains line and column information for each AST node to support debugging

The Parser serves as the critical bridge between the raw text input and the structured AST representation, enabling subsequent components like the Validator and Executor to process GraphQL operations efficiently. By converting the linear token stream into a hierarchical tree structure, the Parser facilitates pattern matching, validation rule application, and execution planning across the GraphQL processing pipeline.

##### Language Location
Location is the component responsible for tracking and maintaining position information within the GraphQL source text during parsing and processing.

The Location component provides precise positional data that enables accurate error reporting and debugging capabilities. It tracks multiple types of position information:
- **Line Numbers**: The line position where tokens, nodes, or errors occur in the source text
- **Column Numbers**: The column position within a specific line
- **Character Offsets**: The absolute character position from the beginning of the source
- **Source References**: Links back to the original source document

The Location information is embedded throughout the parsing pipeline:
- **Tokens**: Each token produced by the Lexer includes location data
- **AST Nodes**: Every AST node contains location information for the source text it represents
- **Errors**: GraphQL errors include precise location data to help developers identify problematic areas
- **Validation**: Location data enables context-aware validation messages

The Location component is essential for developer experience as it transforms generic parsing or validation errors into actionable feedback. Instead of reporting "syntax error," the system can report "syntax error at line 15, column 23," allowing developers to quickly locate and fix issues in their GraphQL queries. This positional tracking is maintained throughout the entire GraphQL processing lifecycle, from initial tokenization through final execution, ensuring that any errors or warnings can be traced back to their precise origin in the source text.

##### Language Visitor
Visitor is the component responsible for traversing and processing Abstract Syntax Tree (AST) nodes using the visitor pattern, providing a systematic and extensible approach to AST manipulation and analysis.

The Visitor component implements the visitor design pattern, which separates algorithms from the object structure on which they operate. This enables the definition of new operations on AST nodes without modifying the node classes themselves, promoting extensibility and maintainability in the GraphQL processing pipeline.

The Visitor provides structured AST traversal capabilities:
- **Node Visitation**: Systematically visits each AST node in a defined order (depth-first traversal)
- **Enter/Leave Hooks**: Provides entry and exit points for each node, allowing pre-processing and post-processing operations
- **Type-Specific Handling**: Enables different processing logic for each AST node type (Document, Field, Fragment, etc.)
- **Traversal Control**: Allows visitors to control traversal flow, including skipping subtrees or early termination
- **Context Preservation**: Maintains traversal context and state throughout the visiting process

The Visitor component supports various traversal patterns:
- **Pre-order Traversal**: Processes parent nodes before their children (enter phase)
- **Post-order Traversal**: Processes parent nodes after their children (leave phase)
- **Conditional Traversal**: Allows selective processing based on node types or conditions
- **Stateful Traversal**: Maintains accumulated state across multiple node visits
- **Multi-pass Traversal**: Enables multiple traversal passes for complex analysis tasks

Key visitor operations include:
- **AST Transformation**: Modifies or replaces AST nodes during traversal
- **Information Collection**: Gathers data from AST nodes for analysis purposes
- **Validation Support**: Enables validation rules to be applied systematically across the AST
- **Code Generation**: Supports code generation workflows that require comprehensive AST analysis
- **Query Analysis**: Facilitates query complexity analysis and optimization planning

The Visitor component integrates with other language components:
- **Parser Integration**: Processes ASTs produced by the Parser component
- **Printer Integration**: Works with the Printer to enable AST-to-text transformation
- **Validator Integration**: Provides the foundation for validation rule application
- **Executor Integration**: Supports execution planning and field resolution analysis

Visitor implementation features:
- **Extensible Design**: New visitor types can be easily added without modifying existing code
- **Composable Patterns**: Multiple visitors can be combined for complex processing workflows
- **Error Handling**: Provides mechanisms for error collection and propagation during traversal
- **Performance Optimization**: Implements efficient traversal algorithms to minimize processing overhead
- **Memory Management**: Manages traversal state efficiently to prevent memory leaks during large AST processing

The Visitor component is fundamental to the GraphQL-Go implementation, enabling modular and extensible AST processing that supports validation, execution planning, introspection, and development tooling throughout the GraphQL processing pipeline.

##### Language Printer
Printer is the component responsible for converting Abstract Syntax Tree (AST) nodes back into their textual GraphQL representation, essentially performing the reverse operation of the Parser component.

The Printer component serves as a bridge between the structured AST representation and human-readable GraphQL syntax. It enables various use cases including query formatting, AST manipulation debugging, and GraphQL document generation from programmatically constructed ASTs.

The Printer performs AST-to-text transformation by traversing AST nodes and generating corresponding GraphQL syntax:
- **Document Printing**: Converts document nodes back to complete GraphQL documents
- **Operation Printing**: Transforms operation definition nodes into query, mutation, or subscription syntax
- **Selection Set Printing**: Converts selection set nodes into properly formatted field selections with proper indentation
- **Field Printing**: Handles field nodes including names, aliases, arguments, and nested selections
- **Fragment Printing**: Processes both named fragment definitions and inline fragment nodes
- **Value Printing**: Converts value nodes back to their literal representations (strings, numbers, booleans, lists, objects)
- **Type Printing**: Transforms type reference nodes into their textual type notation
- **Directive Printing**: Handles directive nodes with their arguments and proper placement

The Printer component works closely with the Visitor component to traverse the AST structure:
- **AST Traversal**: Uses visitor pattern to systematically process each node type
- **Node Transformation**: Applies specific formatting rules for each AST node kind
- **String Building**: Constructs the final GraphQL string representation incrementally
- **Formatting Control**: Maintains proper indentation, spacing, and line breaks for readable output

Key features of the Printer implementation include:
- **Reversible Operations**: Ensures that Parser → Printer → Parser operations preserve semantic meaning
- **Format Consistency**: Produces consistently formatted GraphQL output regardless of input formatting
- **Debugging Support**: Enables developers to inspect and understand AST transformations
- **Integration with Visitor**: Leverages the visitor pattern for extensible and maintainable AST traversal

The Printer component is essential for development tooling, query analysis, and any scenario where programmatically generated or modified ASTs need to be converted back to GraphQL syntax for human consumption or further processing.

##### Executor / Resolver

The Executor/Resolver is a core component responsible for executing GraphQL operations and resolving the fields requested in a query. It forms the bridge between the parsed GraphQL Abstract Syntax Tree (AST) and the actual data retrieval logic, orchestrating the process from query execution to final response construction.

- **Executor**: Traverses the AST produced by the Parser, determining the operation type (query, mutation, subscription) and evaluating each node.
- **Resolver**: Each field in the schema can specify a resolver function. The resolver is called by the Executor to fetch or compute the data for that field. If no custom resolver is specified, a default resolver is used, often retrieving values directly from the source object.

**Responsibilities:**
- Receives an operation (query/mutation/subscription) and variables.
- Traverses the AST, matching fields to resolver functions.
- Calls each resolver with the parent object, arguments, context, and info about the execution state.
- Handles asynchronous and nested field resolution, supporting fragments, directives, and error propagation.
- Aggregates the resolved data into the final response structure.

This design allows for extensibility—custom business logic, authentication, and side effects can be injected at the resolver level. The clear separation also makes the execution engine reusable and testable.

Implementation details can be found throughout the codebase, notably in files and PRs related to Executor and Resolver components.

References:
- [Partial implementation of resolve fields](https://github.com/graphql-go/graphql/pull/8)
- [Field resolver interface support](https://github.com/graphql-go/graphql/pull/288)

##### Validator

The Validator component is responsible for ensuring that incoming GraphQL queries conform to the GraphQL specification and the schema defined by the application. After the Parser creates the Abstract Syntax Tree (AST) from the query, the Validator traverses the AST and applies a set of validation rules.

**Key responsibilities:**
- Confirms that the query structure and field selections are valid according to the schema.
- Ensures that fragments and operations are used correctly and do not conflict.
- Checks for type compatibility, required arguments, valid variable usage, and directive application.
- Leverages the Visitor pattern for AST traversal, enabling modular validation rule implementations.
- Produces detailed and actionable error messages for invalid queries.

The Validator is essential for catching errors and enforcing GraphQL best practices before execution proceeds, contributing to the reliability and security of the system.

Notable improvements include:
- Enhanced error messages with suggestions.
- Deepened introspection queries.
- Support for schema language directives and deprecations.

References:
- [Visitor wired to Printer and Validator](https://github.com/graphql-go/graphql/pull/10)
- [Validation improvements and error handling](https://github.com/graphql-go/graphql/pull/192)

##### Definition

The Definition component is responsible for providing the building blocks and structures needed to define GraphQL schemas programmatically. It contains the core type definitions and interfaces that developers use to construct their GraphQL type system.

**Key responsibilities:**
- Provides foundational type definitions for creating GraphQL schemas including Object types, Interface types, Union types, Enum types, and Scalar types.
- Defines the structure and interfaces for field definitions, argument definitions, and input type definitions.
- Establishes the contracts and behavior patterns that all GraphQL types must follow.
- Enables programmatic schema construction by providing the necessary building blocks and type safety.
- Supports schema extension and composition through modular type definitions.

**Core definition types provided:**
- **ObjectTypeDefinition**: Defines GraphQL object types with fields and resolvers
- **InterfaceTypeDefinition**: Defines GraphQL interface types for polymorphic schemas
- **UnionTypeDefinition**: Defines GraphQL union types for type alternatives
- **EnumTypeDefinition**: Defines GraphQL enum types for predefined value sets
- **ScalarTypeDefinition**: Defines custom GraphQL scalar types
- **InputTypeDefinition**: Defines GraphQL input types for mutations and field arguments
- **FieldDefinition**: Defines individual fields within types including their types, arguments, and resolvers
- **ArgumentDefinition**: Defines arguments for fields and directives

The Definition component serves as the foundation for schema construction, enabling developers to build type-safe GraphQL APIs by providing well-defined interfaces and structures. It ensures consistency across the schema definition process and facilitates schema validation and introspection capabilities.

This component integrates closely with the Schema component to construct complete GraphQL schemas and with the Validator component to ensure schema definitions conform to GraphQL specification requirements.

##### Schema

The Schema component is the central definition and organizational structure for all types, queries, mutations, and subscriptions in a GraphQL service. It acts as the contract between the client and server, specifying what operations and data are available, and how they relate.

**Key responsibilities:**
- Defines root operation types (Query, Mutation, Subscription).
- Registers all object types, scalars, enums, interfaces, unions, and input types.
- Describes relationships between types, fields, arguments, and return values.
- Provides the basis for the Validator to enforce query correctness and for the Executor to resolve fields.
- Supports schema introspection, allowing clients to query the schema structure itself.
- Integrates schema language features (e.g., directives, deprecations, custom extensions).

The Schema enables modular, type-safe API development, supporting both code-first and schema-first approaches. In GraphQL-Go, schema construction can be programmatic, mirroring the graphql-js reference implementation, with Go types and resolvers mapped to schema elements.

Implementation enhancements include:
- Support for schema language directives and introspection improvements.
- Schema definition integration into the language components.
- Enhanced handling of types and arguments, including context propagation to executors.

References:
- [Schema language directives & improvements](https://github.com/graphql-go/graphql/pull/123)
- [Schema validation and introspection enhancements](https://github.com/graphql-go/graphql/pull/192)

##### Directives

The Directives component manages the definition, registration, and application of directives in the GraphQL execution pipeline. Directives are special annotations that can be attached to fields, fragments, and other schema elements to modify query behavior at runtime or affect schema interpretation.

**Key responsibilities:**
- Defines built-in directives (such as `@include`, `@skip`, and `@deprecated`) and supports user-defined custom directives.
- Registers directive locations (e.g., on fields, fragments, schema definitions) in compliance with the GraphQL specification.
- Applies directive logic during query validation and execution phases, allowing conditional field inclusion/exclusion, deprecation warnings, and extensible custom behaviors.
- Ensures that directives are type-safe, context-aware, and validated against their declared arguments and locations.
- Supports directive introspection, enabling clients to discover available directives and their intended use.

Directives extend the expressiveness and adaptability of the GraphQL schema and queries, supporting common patterns like conditional data fetching and schema evolution with minimal changes.

Implementation highlights:
- Integrates tightly with the schema, validator, and executor components.
- Enables extensible directive creation for reusable, cross-cutting concerns.

References:
- [Schema language directives & improvements](https://github.com/graphql-go/graphql/pull/123)
- [Schema validation and introspection enhancements](https://github.com/graphql-go/graphql/pull/192)

##### Rules

The Rules component encapsulates the set of validation rules applied to GraphQL operations to ensure they conform to the GraphQL specification and the schema’s constraints. Each rule encapsulates a specific aspect of query validity, such as detecting field selection conflicts, enforcing fragment naming, or checking variable usage.

**Key responsibilities:**
- Defines and organizes individual validation rules, each targeting a different aspect of query or schema correctness according to the GraphQL specification.
- Integrates with the Validator and Visitor components to traverse the AST and apply rules systematically.
- Reports detailed error messages for violations, aiding developers in quickly resolving issues.
- Supports extensibility, allowing custom rules to be defined and composed with built-in rules.
- Handles advanced concerns such as overlapping fields, argument defaults, and fragment cycles.

The Rules component ensures the reliability and safety of GraphQL queries by systematically enforcing correctness, enabling robust validation workflows for both built-in and user-defined constraints.

Implementation highlights:
- Maintains parity with the reference graphql-js implementation.
- Improvements include enhanced validation error messages and deeper introspection for more comprehensive rule enforcement.

References:
- [Schema validation and error message improvements](https://github.com/graphql-go/graphql/pull/192)
- [Rules and validation enhancements](https://github.com/graphql-go/graphql/pull/123)

##### Introspection

The Introspection component enables clients and tools to query the schema structure and metadata at runtime. This functionality is central to GraphQL’s self-describing nature, supporting features like schema documentation, auto-completion, validation, and IDE integrations.

**Key responsibilities:**
- Implements the introspection system as defined by the GraphQL specification, exposing types, fields, directives, and their relationships via special introspection queries (such as `__schema`, `__type`, and `__directive`).
- Integrates with the Executor and Schema components to resolve introspection fields dynamically based on the current schema.
- Supports deep, recursive queries to inspect nested types and directives, enabling clients to fully explore the schema capabilities.
- Provides compatibility with tools and environments that rely on introspection, such as GraphQL Playground, GraphiQL, and code generators.

Implementation details:
- The introspection system is kept up to date with changes from the graphql-js reference, ensuring compatibility and feature parity.
- The standard introspection query used for validation and testing can be found at:
  [testutil/introspection_query.go](https://github.com/graphql-go/graphql/blob/58689e0742f217137ac0c82b16e2d658e0cc1853/testutil/introspection_query.go#L4)

References:
- [GraphQL Introspection Query (Go implementation)](https://github.com/graphql-go/graphql/blob/58689e0742f217137ac0c82b16e2d658e0cc1853/testutil/introspection_query.go#L4)

##### Language/TypeInfo

The `TypeInfo` component provides utilities for tracking and managing GraphQL type information during syntax tree traversal, such as when validating or analyzing queries. It encapsulates the logic required to keep track of the current type context, input and output types, field definitions, directives, and arguments as the AST is visited.

**Key responsibilities:**
- Maintains the current type context while traversing the GraphQL AST (Abstract Syntax Tree), updating its internal state as new nodes are entered or exited.
- Exposes methods to access the current type, parent type, input type, field definition, directive, and argument.
- Supports validation and schema analysis by providing accurate and up-to-date type information at any point in the traversal.
- Integrates with the Validator and Visitor components, serving as the backbone for context-aware validation rules and advanced analysis.
- Enables correct interpretation of fragments, fields, arguments, and directives as specified in the GraphQL specification.

**Implementation highlights:**
- The design mirrors the reference implementation in graphql-js, ensuring feature parity and reliability.
- TypeInfo is stateful and designed for reuse across different traversals within the same schema.
- Provides a clear interface for external tools and rules to query type context without directly mutating internal state.

References:
- [GraphQL Specification: Type System](https://spec.graphql.org/June2018/#sec-Type-System)
- [graphql-js TypeInfo source](https://github.com/graphql/graphql-js/blob/main/src/utilities/TypeInfo.ts)

##### Compatibility

The compatibility-unit-tests framework contains the following main components that enable validation of GraphQL implementation compatibility against the graphql-js reference implementation:

- **App**: The main application orchestrator that coordinates the compatibility validation process between different GraphQL implementations.

- **Cmd**: The command-line interface component that provides user interaction capabilities for selecting GraphQL implementations and configuring validation parameters.

- **Extractor**: The component responsible for extracting unit test names and structures from GraphQL implementations, supporting both JavaScript (graphql-js) and Go implementations.

- **Implementation**: The component that defines and manages different GraphQL implementation configurations, including repository information, test extraction strategies, and implementation-specific metadata.

- **Result**: The component that handles the collection, aggregation, and presentation of compatibility validation results, including success/failure metrics and detailed comparison outcomes.

- **Types**: The component that defines the core data structures and type definitions used throughout the compatibility validation framework.

- **Validator**: The component that performs the actual compatibility comparison between the reference implementation (graphql-js) and target implementations, validating test coverage and identifying compatibility gaps.

The compatibility framework operates by extracting unit test names from both the reference graphql-js implementation and the target implementation, then comparing them to identify missing tests or implementation gaps. This systematic approach ensures that GraphQL implementations maintain compatibility with the standard reference implementation.

References:
- [Compatibility Unit Tests Repository](https://github.com/graphql-go/compatibility-unit-tests)
- [Compatibility validation framework documentation](https://github.com/graphql-go/compatibility-unit-tests#validation-compatibility)

##### Standard Definitions

The compatibility-standard-definitions framework contains the following main components that enable validation of GraphQL implementation type system compatibility against the GraphQL specification and the graphql-js reference implementation:

- **Puller**: The component responsible for cloning GraphQL repositories including the GraphQL specification repository, the GraphQL JavaScript reference implementation repository, and target GraphQL implementation repositories. It manages the retrieval and synchronization of source code from multiple repositories to enable cross-implementation compatibility analysis.

- **Extractor**: The component that extracts type system definitions from different sources using multiple strategies. It pulls type system definitions by parsing from the GraphQL specification repository, by introspection from the GraphQL JavaScript reference implementation, and by introspection from target GraphQL implementations. This component bridges different extraction methodologies to create a unified comparison baseline.

- **Executor**: The component that executes type system definitions introspection results on GraphQL implementations. It processes the extracted type system definitions and applies them to target implementations to validate their compatibility and correctness against the standard definitions.

- **Validator**: The component that performs comprehensive validation by comparing schemas and type system definitions across different implementations. It validates the GraphQL specification schema against the GraphQL JavaScript reference implementation schema, compares GraphQL implementation schemas against the GraphQL specification schema, and performs detailed comparisons of type system definitions between the specification and implementations.

The standard definitions framework operates by systematically extracting, processing, and comparing type system definitions from the GraphQL specification, the reference implementation, and target implementations. This approach ensures that GraphQL implementations maintain structural compatibility with the standard type system definitions as defined by the GraphQL specification.

References:
- [Compatibility Standard Definitions Repository](https://github.com/graphql-go/compatibility-standard-definitions)
- [GraphQL specification repository validation](https://github.com/graphql-go/compatibility-standard-definitions#implementation-details)

##### User Acceptance Definitions

The compatibility-user-acceptance framework contains the following main components that enable validation of GraphQL implementation developer experience metrics against the GraphQL specification and the graphql-js reference implementation:

- **Extractor**: The component responsible for extracting GitHub repository metrics from GraphQL implementations. It fetches comprehensive repository data including stars count, issues statistics, pull requests metrics, fork counts, license information, last commit dates, contributor counts, and GraphQL specification version compatibility. The extractor provides both live GitHub API integration and test data capabilities for development and validation purposes.

- **Repository**: The component that contains extracted repository metrics as structured data. It serves as the data model for storing and organizing the GitHub repository information collected by the extractor, providing a standardized interface for accessing metrics such as stars count, issues, pull requests, and other repository characteristics used in compatibility validation.

- **CLI Tool**: The component that provides a command-line interface for running compatibility validation between GraphQL implementations. It orchestrates the extraction process, compares metrics against reference thresholds, calculates difference ratios, and presents results in a user-friendly table format. The CLI tool validates compatibility by comparing various developer experience metrics including repository engagement, community activity, and specification compliance indicators.

The user acceptance framework operates by systematically extracting GitHub repository metrics from target GraphQL implementations, comparing them against the graphql-js reference implementation, and calculating difference ratios to determine compatibility levels. This approach ensures that GraphQL implementations maintain acceptable developer experience standards and community engagement metrics that align with the reference implementation's user acceptance patterns.

References:
- [Compatibility User Acceptance Repository](https://github.com/graphql-go/compatibility-user-acceptance)
- [GraphQL implementation user acceptance validation](https://github.com/graphql-go/compatibility-user-acceptance#use-cases)

### 4.10 Architecture

#### UML Packages Diagram

<img width="817" height="939" alt="Image" src="https://raw.githubusercontent.com/chris-ramon/thesis-graphql-go/refs/heads/main/architecture/uml/packages/graphql-graphql-go-architecture-packages.drawio.png" />

#### UML Components Diagram

<img width="817" height="939" alt="Image" src="https://raw.githubusercontent.com/chris-ramon/thesis-graphql-go/refs/heads/main/architecture/uml/components/graphql-graphql-go-architecture-components.drawio.png" />

#### UML Activity Diagram

<img width="817" height="939" alt="Image" src="https://raw.githubusercontent.com/chris-ramon/thesis-graphql-go/refs/heads/main/architecture/uml/activity/graphql-graphql-go-architecture-activity.drawio.png" />

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

Go open-source library that implements the `graphql-js` reference implementation.


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

Go open-source library that obtains GitHub information.


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
